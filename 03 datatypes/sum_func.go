package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(sum(os.Args[1], os.Args[2]))
}

func sum(n1 string, n2 string) (result int) {
	n1_i, _ := strconv.Atoi(n1)
	n2_i, _ := strconv.Atoi(n2)
	result = n1_i + n2_i
	return
}
