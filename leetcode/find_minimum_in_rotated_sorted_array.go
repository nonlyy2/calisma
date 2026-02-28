package leetcode

import "fmt"

// Задача №153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for right > left {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
			fmt.Printf("теперь лефт(%d) равен мид(%d)\n", left, mid)
			continue
		} else {
			right = mid
			fmt.Printf("теперь райт(%d) равен мид(%d)\n", right, mid)

		}
	}

	return nums[left]
}
