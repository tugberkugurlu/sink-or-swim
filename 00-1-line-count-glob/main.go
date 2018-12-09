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
	var wg sync.WaitGroup

	// see https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
	for _, match := range matches {
		wg.Add(1)
		go func (filePath string) {
			defer wg.Done()
			countOfLines, err := filelinecounter.CountLines(filePath, false)
			if err == nil {
				ch <- countOfLines
			}
		}(match)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for m := range ch {
		fmt.Println(m)
	}
}
