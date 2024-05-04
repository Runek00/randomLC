package main

import (
	"fmt"
	"slices"
)

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	boats := 0
	left := 0
	right := len(people) - 1
	for left <= right {
		boats++
		if people[right]+people[left] <= limit {
			left++
		}
		right--
	}
	return boats
}
func main() {
	fmt.Println(numRescueBoats([]int{1, 2}, 3))
	fmt.Println(numRescueBoats([]int{3, 2, 2, 1}, 3))
	fmt.Println(numRescueBoats([]int{3, 5, 3, 4}, 5))

}
