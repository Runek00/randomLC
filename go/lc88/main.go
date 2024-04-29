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
	for j, k := range nums2 {
		for k > nums1[i] && i < m + j {
			i = i + 1
		}
		nums1 = slices.Insert(nums1, i, k)
	}
	nums1 = nums1[0:m+n]
}
