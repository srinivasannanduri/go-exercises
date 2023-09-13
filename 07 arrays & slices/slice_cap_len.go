package main

import (
	"fmt"
)

func main() {
	var arr [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 89, 90, 100}

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	slice := arr[1:8]

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[6] = 80
	fmt.Println(arr)
	fmt.Println(slice)

	// Original array doesn't change here
	slice = append(slice, 900, 1000, 2000)
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
