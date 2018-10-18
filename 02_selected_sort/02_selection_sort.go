package main

func selectionSort(origin []int) []int {
	for i := 0; i < len(origin); i++ {
		for j := i; j < len(origin); j++ {
			if origin[i] > origin[j] {
				origin[i], origin[j] = origin[j], origin[i]
			}
		}
	}
	return origin
}
