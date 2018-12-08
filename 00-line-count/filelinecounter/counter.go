package filelinecounter

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
)

// CountLines returns number of the lines in a given file. It skips the empty lines
// if ignoreEmptyLines is provided.
func CountLines(absoluteFilePath string, ignoreEmptyLines bool) (count int64, err error) {
	isDir, _ := isDirectory(absoluteFilePath)
	if isDir {
		err = errors.New("path is a directory, not a file")
		return
	}

	// https://groups.google.com/forum/#!topic/golang-nuts/_6YqjJdfYyA
	if isExecutable(absoluteFilePath) == nil {
		// https://golang.org/pkg/errors/
		err = errors.New("the file is executable, skipping the count check")
		return
	}

	file, e := os.Open(absoluteFilePath)
	if e != nil {
		err = e
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ignoreEmptyLines {
			line := scanner.Text()
			if len(line) == 0 {
				continue
			}
		}
		count++
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

// isExecutable returns an error if a given file is not an executable.
// see https://golang.org/src/os/executable_path.go for more info
func isExecutable(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	mode := stat.Mode()
	if !mode.IsRegular() {
		return errors.New("Permission error")
	}
	if (mode & 0111) == 0 {
		return errors.New("Permission error")
	}
	return nil
}
