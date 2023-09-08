package main

import (
	"fmt"
)

func main() {
	var arr [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 89, 90, 100}

	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	slice := arr[1:8]
	fmt.Println(len(slice))

	fmt.Println(slice)
	fmt.Println(cap(slice))
}
