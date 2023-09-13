package main

import "fmt"

func main() {
	// Delete element at index 2

	arr := [5]int{10, 20, 30, 40, 50}
	i := 2

	slice_1 := arr[:i]
	slice_2 := arr[i+1:]

	result := append(slice_1, slice_2...)
	fmt.Println(arr)
	fmt.Println(result)
}
