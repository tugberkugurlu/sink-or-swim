package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	// sender
	go func() {
		count := 0
		for {
			count++
			fmt.Println("<- Sending ", count)
			ch1 <-count
			fmt.Println("<- Sent ", count)
		}
	}()

	// receiver
	go func() {
		for item := range ch1 {
			fmt.Println("    -> processing", item)
			time.Sleep(2 * time.Second)
			fmt.Println("    -> processed", item)
		}
	}()

	// block forever without "eating" up your CPU
	// https://stackoverflow.com/a/36419222/463785
	select{}
}
