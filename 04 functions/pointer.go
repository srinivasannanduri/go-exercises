package main

import "fmt"

func main() {
	var name string = "Srinu"
	updateName(&name)

	fmt.Println(name)

}

func updateName(name *string) {
	*name = "Srinivasan"
}
