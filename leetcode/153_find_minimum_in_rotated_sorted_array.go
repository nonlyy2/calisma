package leetcode

// Задача №153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for right > left {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
			continue
		} else {
			right = mid
		}
	}

	return nums[left]
}
