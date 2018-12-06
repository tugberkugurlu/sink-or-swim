package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"./filelinecounter"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func main() {
	path := flag.String("path", "", "Path to the file to count its lines")
	ignoreEmptyLines := flag.Bool("ignore-empty-lines", false, "Specifies to ignore empty lines")

	flag.Parse()

	if *path == "" {
		flag.Usage()
		os.Exit(1)
	}

	absoluteFilePath, _ := filepath.Abs(*path)
	doesExist, _ := exists(absoluteFilePath)
	if !doesExist {
		fmt.Println("path does not exist", absoluteFilePath)
		flag.Usage()
		os.Exit(1)
	}

	isDir, _ := isDirectory(absoluteFilePath)
	if isDir {
		fmt.Println("path is a directory, not a file", absoluteFilePath)
		flag.Usage()
		os.Exit(1)
	}

	count := filelinecounter.CountLines(absoluteFilePath, *ignoreEmptyLines)

	fmt.Println("==========================")
	fmt.Println("Count: ", count)
	fmt.Println("==========================")

	// fmt.Println(absoluteFilePath)
	// fmt.Println("hello world")
	// fmt.Println(os.Args[1:])
	// fmt.Println(*path)
	// fmt.Println(*ignoreEmptyLines)
}
