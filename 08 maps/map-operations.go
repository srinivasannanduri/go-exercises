package main

import "fmt"

func main() {
	// codes := make(map[string]int)
	// codes["A"] = 65
	// codes["F"] = 70
	// codes["K"] = 75

	codes := map[string]int{
		"A": 65, "F": 70, "K": 75,
	}

	delete(codes, "F")

	fmt.Println(codes)

}
