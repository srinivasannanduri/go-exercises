package main

import (
	"fmt"
	"os"
	"strconv"
)

func fibonacci(n int) []int {
	if n < 2 {
		return make([]int, 0)
	}

	nums := make([]int, n)
	nums[0], nums[1] = 1, 1

	for i := 2; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}

func main() {
	target, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("The Fibonacci sequence is:", fibonacci(target))
}
