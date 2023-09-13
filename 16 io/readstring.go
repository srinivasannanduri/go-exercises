package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Let us catch up over a cup of coffee")

	buf := make([]byte, 5)
	for {
		n, err := reader.Read(buf)
		fmt.Println(buf[:n], err)
		if err != nil {
			break
		}
	}
}
