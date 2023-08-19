package main

import "fmt"

func number() int {
	return 1
}

func main() {
	if num := number(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is positive with single digit")
	} else {
		fmt.Println(num, "is positive with multiple digits")
	}
}
