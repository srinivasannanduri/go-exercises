package main

import (
	"fmt"
)

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	for name, _ := range studentsAge {
		fmt.Printf("Names %s\n", name)

		// fmt.Printf("Ages %d\n", age)

	}
}
