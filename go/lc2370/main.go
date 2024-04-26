package main

import "fmt"

func longestIdealString(s string, k int) int {
	lastchar := make(map[int]int, 0)
	for _, c := range s{
		cur := int(c)
		curLength := 0
		for j := -k; j <= k; j++ {
			lastLength, ok := lastchar[cur + j]
			if ok {
				if lastLength > curLength {
					curLength = lastLength
				}
			}
		}
		curLength = curLength + 1
		lastchar[cur] = curLength
	}
	out := 0
	for _, l := range lastchar {
		if l > out {
			out = l
		}
	}
	return out
}

func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
	fmt.Println(longestIdealString("abcd", 3))
}
