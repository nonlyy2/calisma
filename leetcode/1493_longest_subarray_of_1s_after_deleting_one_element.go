package leetcode

// Задача №1493: Longest Subarray of 1's After Deleting One Element
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/

func LongestSubarray(nums []int) int {
	maxLen := 0
	zeroCount := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		if zeroCount > 1 {
			for nums[left] == 1 {
				left++
			}
			left++
			zeroCount--
		}

		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
