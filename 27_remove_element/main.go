package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 3), nums)
	fmt.Println(removeElement(nums2, 2), nums2)
}

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
