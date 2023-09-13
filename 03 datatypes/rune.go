package main

import "fmt"

func main() {
	str := "String"
	resultRune := make([]rune, len(str)*2)
	chars := []rune(str)

	/*currentResultRuneIndex := 0
	for _, char := range chars {
		resultRune[currentResultRuneIndex] = char
		resultRune[currentResultRuneIndex+1] = char
		currentResultRuneIndex = currentResultRuneIndex + 2
	}*/

	for i, char := range chars {
		idx := i * 2
		resultRune[idx] = char
		resultRune[idx+1] = char
	}

	fmt.Println(string(resultRune))
}
