package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		for {
			select {
			case firstValue, ok := <-firstChan:
				if !ok {
					return
				}
				resultChan <- firstValue * firstValue
				return

			case secondValue, ok := <-secondChan:
				if !ok {
					return
				}
				resultChan <- secondValue * 3
				return

			case <-stopChan:
				return
			}
		}
	}()

	return resultChan
}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	stopChan := make(chan struct{})

	go func() {
		firstChan <- 5
	}()

	go func() {
		secondChan <- 10
	}()

	go func() {
		stopChan <- struct{}{}
	}()

	resultChan := calculator(firstChan, secondChan, stopChan)

	for result := range resultChan {
		fmt.Println(result)
	}
	close(stopChan)
}
