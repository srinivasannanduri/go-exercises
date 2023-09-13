package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Create first goroutine here
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
	// Create second goroutine here

	// Wait for both goroutines to finish here
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {
	for i := 'a'; i < 'f'; i++ {
		fmt.Println(string(i))
	}

	wg.Done()
}
