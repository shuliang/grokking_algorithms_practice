package main

import (
	"fmt"
)

const approve = "let them vote!"
const reject = "kick them out!"

var voted = map[string]bool{}

func check(name string) string {
	if _, ok := voted[name]; ok {
		return reject
	}
	voted[name] = true
	return approve
}

func main() {
	fmt.Println(check("tom"))
	fmt.Println(check("mike"))
	fmt.Println(check("mike"))
}
