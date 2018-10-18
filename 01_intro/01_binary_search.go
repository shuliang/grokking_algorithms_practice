package main

import (
	"fmt"
)

// list must be acsend ordered list
func binarySearch(list []int, target int) int {
	start := 0
	end := len(list) - 1
	for start <= end {
		mid := (start + end) / 2
		// fmt.Println("------")
		// fmt.Println("start:", start)
		// fmt.Println("end:  ", end)
		// fmt.Println("nMid: ", mid)
		if list[mid] == target {
			return mid
		}
		if list[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("found %d in %d\n", 4, binarySearch(list, 4))
	fmt.Println("****************")

	list2 := []int{1, 3, 5, 7, 9}
	fmt.Printf("found %d in %d\n", 3, binarySearch(list2, 3))
	fmt.Println("****************")
	fmt.Printf("found %d in %d\n", 1, binarySearch(list2, 1))
	fmt.Println("****************")
	fmt.Printf("found %d in %d\n", 9, binarySearch(list2, 9))
	fmt.Println("****************")
	fmt.Printf("found %d in %d\n", -1, binarySearch(list2, -1))
}
