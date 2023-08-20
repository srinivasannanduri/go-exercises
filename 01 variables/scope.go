package main

import "fmt"

var name string = "Sri"

func main() {
	fmt.Println(name)

	var country string = "India"

	{
		var state string = "Karnataka"
		fmt.Println(state)
		fmt.Println(country)
		// fmt.Println(place) -- Doesn't work
	}

	// var place string = "Bengaluru"
	// fmt.Println(state) -- Doesn't work
}
