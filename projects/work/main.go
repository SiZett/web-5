package main

import (
	"fmt"
	"time"
	// "sync"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
	defer close(outputStream)

	prevValue := <-inputStream
	outputStream <- prevValue

	for value := range inputStream {
		if prevValue == value {
			continue
		}

		outputStream <- value
		prevValue = value
	}
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go func() {
		defer close(inputStream)
		inputStream <- "Hello"
		inputStream <- "World"
		inputStream <- "World"
		inputStream <- "!"
	}()

	go removeDuplicates(inputStream, outputStream)

	for value := range outputStream {
		fmt.Println(value)
	}
}
