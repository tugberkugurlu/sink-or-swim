package main

import (
	"bufio"
	"os"
)

// CountLines returns number of the lines in a given file. It skips the empty lines
// if ignoreEmptyLines is provided.
func CountLines(absoluteFilePath string, ignoreEmptyLines bool) (count int64) {
	file, err := os.Open(absoluteFilePath)
	if err != nil {
		panic(err)
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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}
