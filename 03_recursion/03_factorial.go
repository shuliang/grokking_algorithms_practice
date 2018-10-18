package main

import (
	"fmt"
)

func factorial(input int) int {
	if input < 0 {
		return -1
	} else if input == 1 {
		return 1
	}
	return input * factorial(input-1)
}

func main() {
	fmt.Println(factorial(5))
}
