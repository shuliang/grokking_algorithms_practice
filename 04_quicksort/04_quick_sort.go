package main

import (
	"fmt"
)

func loopSum(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return total
}

func recursiveSum(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + recursiveSum(list[1:])
}

func count(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + count(list[1:])
}

func findMax(list []int) int {
	if len(list) == 2 {
		if list[0] >= list[1] {
			return list[0]
		}
		return list[1]
	}

	subMax := findMax(list[1:])
	if list[0] >= subMax {
		return list[0]
	}
	return subMax
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	mid := len(list) / 2
	pivot := list[mid]
	left, right := []int{}, []int{}
	for i, v := range list {
		if i != mid {
			if pivot > v {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	result := append(left, pivot)
	result = append(result, right...)
	return result
}

func qSortPivotFirst(list []int) []int {
	if len(list) < 2 {
		return list
	}
	left, right := []int{}, []int{}
	pivot := list[0]
	for _, v := range list[1:] {
		if pivot > v {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = qSortPivotFirst(left)
	right = qSortPivotFirst(right)
	all := append(left, pivot)
	all = append(all, right...)
	return all
}

func main() {
	list := []int{1, 29, 20, 38, 14, 5, 72, 3, 3, 2}
	fmt.Printf("loop sum of %v is %d\n", list, loopSum(list))
	fmt.Printf("recu sum of %v is %d\n", list, recursiveSum(list))
	fmt.Printf("li count of %v is %d\n", list, count(list))
	fmt.Printf("find max in %v is %d\n", list, findMax(list))
	fmt.Printf("quick sort  %v is %d\n", list, quickSort(list))
	fmt.Printf("quick sort  %v is %d\n", list, qSortPivotFirst(list))
}
