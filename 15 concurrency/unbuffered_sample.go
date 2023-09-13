package main

import (
	"fmt"
	"time"
)

func main() {
	var items chan string
	items = make(chan string)
	go sell(items)
	go buy(items)
	time.Sleep(2 * time.Second)
}

func sell(items chan string) {
	fmt.Println("Selling Furniture")
	items <- "Furniture"
}

func buy(items chan string) {
	fmt.Println("Waiting for items to be sold")
	item, ok := <-items
	fmt.Println("Received item ", item, ok)
}
