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

	count, err := filelinecounter.CountLines(absoluteFilePath, *ignoreEmptyLines)

	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Count: ", count)
		fmt.Println("==========================")
	} else {
		fmt.Println("==========================")
		fmt.Println("Error: ", err)
		fmt.Println("==========================")
	}

	// fmt.Println(absoluteFilePath)
	// fmt.Println("hello world")
	// fmt.Println(os.Args[1:])
	// fmt.Println(*path)
	// fmt.Println(*ignoreEmptyLines)
}
