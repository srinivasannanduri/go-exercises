package main

import "fmt"

func main() {
	for i := 2; i <= 20; i++ {
		if isPrime(i) {
			fmt.Printf("%v ", i)
		}
	}
}

func isPrime(num int) bool {
	for j := 2; j <= num/2; j++ {
		if num%j == 0 {
			return false
		}
	}

	return true
}
