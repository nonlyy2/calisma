package leetcode

// Задача №3105: Longest Strictly Increasing or Strictly Decreasing Subarray
// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/

func LongestMonotonicSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	incLen := 1
	decLen := 1
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			incLen++
			decLen = 1
		} else if nums[i] < nums[i-1] {
			decLen++
			incLen = 1
		} else {
			incLen = 1
			decLen = 1
		}

		maxLen = max(maxLen, incLen, decLen)
	}

	return maxLen
}
