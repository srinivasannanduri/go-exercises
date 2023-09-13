package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// your code goes here
	file, err := os.Open("/Users/sn/work/code/srinivasannanduri/go-exercises/16 io/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	sum := 0
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}

		sum += num
	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println(err)
	}

	out, err := os.Create("/Users/sn/work/code/srinivasannanduri/go-exercises/16 io/output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer out.Close()

	_, err = out.WriteString(strconv.Itoa(sum))
	if err != nil {
		fmt.Println(err)
		return
	}
}
