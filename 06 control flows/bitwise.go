package main

import "fmt"

func main() {
	var x, y int = 100, 90
	fmt.Println(x & y)
	fmt.Println(x | y)
}

// 100 - 1100100
//  90 - 1011010
//       -------
//       1000000 - & = 64
// 	  -------
// 	  -------
//       1111110 - | = 126
// 	  -------
