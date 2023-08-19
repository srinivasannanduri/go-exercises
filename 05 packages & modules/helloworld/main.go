package main

import (
	"fmt"

	"github.com/srinivasannanduri/calculator"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(calculator.Sum(20, 30))

	fmt.Println(quote.Hello())

	// total := calculator.internalSum(5)
	// fmt.Println(total)
	// fmt.Println("Version: ", calculator.logMessage)
}
