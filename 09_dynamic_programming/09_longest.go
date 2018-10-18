package main

import (
	"fmt"
)

func longestCommonSubstring(a, b string) string {
	return longest(a, b, false)
}

func longestCommonSubsequence(a, b string) string {
	return longest(a, b, true)
}

// longest if `isSubsequence = true`, return the longest common subsequence,
// otherwise return the longest common substring
func longest(a, b string, isSubsequence bool) string {
	if len(a) == 0 || len(b) == 0 {
		return ""
	}
	runeA := []rune(a)
	runeB := []rune(b)
	cell := make([][]int, len(runeA)*len(runeB))
	for i := range cell {
		cell[i] = make([]int, len(runeB))
	}

	subEndIndex := 0
	longest := 0
	for i, aa := range runeA {
		for j, bb := range runeB {
			if aa == bb {
				if i > 0 && j > 0 {
					cell[i][j] = cell[i-1][j-1] + 1
				} else {
					cell[i][j] = 1
				}
				if cell[i][j] > longest {
					longest = cell[i][j]
					subEndIndex = i
				}
			} else {
				if i > 0 && j > 0 && isSubsequence {
					cell[i][j] = max(cell[i-1][j], cell[i][j-1])
				} else {
					cell[i][j] = 0
				}
			}
		}
	}
	sub := string(runeA[subEndIndex-longest+1 : subEndIndex+1])
	return sub
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println("longests subsequence:", longestCommonSubsequence("fish", "fosh"))
	fmt.Println("longests subsequence:", longestCommonSubsequence("Hello世界7世界", "hello世界7"))
	fmt.Println("longests subsequence:", longestCommonSubsequence("Hello世界", "hello界"))

	fmt.Println("longests substring:  ", longestCommonSubstring("fish", "fosh"))
	fmt.Println("longests substring:  ", longestCommonSubstring("Hello世界7世界", "hello世界7"))
	fmt.Println("longests substring:  ", longestCommonSubstring("Hello世界", "hello界"))
}
