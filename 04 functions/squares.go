package main

import (
	"fmt"
	"strings"
)

func calcSquare(numbers []int) []int {
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
		// fmt.Println(len(squares))
		// fmt.Println(cap(squares))
	}
	return squares

}

func main() {
	nums := [3]int{10, 20, 15}
	fmt.Println(calcSquare(nums[:]))

	result := printStrings("Joe", "Monica", "Gunther")
	fmt.Println(result)
}

func printStrings(names ...string) (names_c []string) {
	names_c = []string{}
	for _, value := range names {
		names_c = append(names_c, strings.ToUpper(value))
	}
	return
}
