package main

import (
	"fmt"
	"strings"
)

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}
