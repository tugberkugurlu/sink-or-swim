package main

import (
	"../00-line-count/filelinecounter"
	"fmt"
	"path/filepath"
	"sync"
)

func main() {
	// get the list of files per glob pattern
	// iterate the list of files
	// calculate the count per each

	absolutePath, _ := filepath.Abs("../../../**/*.*")
	matches, _ := filepath.Glob(absolutePath)
	ch := make(chan int64, len(matches))

	// oh well ¯\_(ツ)_/¯ https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
	for _, match := range matches {
		go func () {
			countOfLines, err := filelinecounter.CountLines(match, false)
			if err == nil {
				ch <- countOfLines
			}
		}()
	}

	for countOfLines := range ch {
		fmt.Println(countOfLines)
	}
}
