package main

import (
	"../00-line-count/filelinecounter"
	"fmt"
	"path/filepath"
)

func main() {
	// get the list of files per glob pattern
	// iterate the list of files
	// calculate the count per each

	absolutePath, _ := filepath.Abs("../../../**/*.*")
	matches, _ := filepath.Glob(absolutePath)

	for _, match := range matches {
		countOfLines, err := filelinecounter.CountLines(match, false)
		if err == nil {
			fmt.Println(match, countOfLines)
		}
	}
}
