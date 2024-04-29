package main

import (
	"fmt"
	"slices"
)

func main() {
	merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := 0
	nums3 := nums1
	for j, k := range nums2 {
		for k > nums3[i] && i < m + j {
			i = i + 1
		}
		nums3 = slices.Insert(nums3, i, k)
	}
	nums1 = nums3[0:m+n]
	fmt.Println(nums1)
}
