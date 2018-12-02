package main

import (
	"path/filepath"
	"fmt"
	"os"
	"flag"
)

func main() {

	path := flag.String("path", "", "Path to the file to count its lines")
	ignoreEmptyLines := flag.Bool("ignore-empty-lines", false, "Specifies to ignore empty lines")

	flag.Parse()

	if *path == "" {
		flag.Usage()
		os.Exit(1)
	}

	absoluteFilePath, _ := filepath.Abs(*path)

	fmt.Println(absoluteFilePath)
	fmt.Println("hello world")
	fmt.Println(os.Args[1:])
	fmt.Println(*path)
	fmt.Println(*ignoreEmptyLines)
}
