package main

import (
	"fmt"
)

func main() {
	var nums1 []int = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var nums2 []int = []int{1, 1, 2}
	var result1 int = removeDuplicates(nums1)
	fmt.Println(result1)
	var result2 int = removeDuplicates(nums2)
	fmt.Println(result2)
}

func removeDuplicates(nums []int) int {
	var duplicateCount int = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[duplicateCount] = nums[i]
			duplicateCount++
		}
	}
	return duplicateCount
}
