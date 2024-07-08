package main

import (
	"fmt"
	"time"
)

func task1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Result from Task 1"
}

func task2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Result from Task 2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	select {
	case result := <-ch1:
		fmt.Println(result)
	case result := <-ch2:
		fmt.Println(result)
	}
}
