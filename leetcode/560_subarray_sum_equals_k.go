package leetcode

// Задача №560: Subarray Sum Equals K
// https://leetcode.com/problems/subarray-sum-equals-k/

func SubarraySum(nums []int, k int) int {
	prefixSums := make(map[int]int)
	prefixSums[0] = 1
	count := 0
	currentSum := 0

	for i := range nums {
		currentSum += nums[i]

		if val, ok := prefixSums[currentSum-k]; ok {
			count += val
		}

		prefixSums[currentSum]++
	}

	return count
}
