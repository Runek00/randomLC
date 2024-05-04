package main

import (
	"fmt"
	"slices"
)

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	freq := make(map[int]int)
	for _, v := range people {
		freq[v]++
	}
	boats := 0
	for i := len(people) - 1; i >= 0; i-- {
		if freq[people[i]] == 0 {
			continue
		}
		if people[i] == limit {
			freq[people[i]]--
			boats++
			continue
		}
		j := limit - people[i]
		freq[people[i]]--
		for j >= 1 {
			if freq[j] > 0 {
				freq[j]--
				break
			}
			j--
		}
		boats++
	}
	return boats
}
func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))

}
