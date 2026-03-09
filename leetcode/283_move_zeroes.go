package leetcode

// Задача №283: Move Zeroes
// https://leetcode.com/problems/move-zeroes/

func MoveZeroes(nums []int) {
	nonZeroCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroCount] = nums[i]
			nonZeroCount++
		}
	}

	for i := nonZeroCount; i < len(nums); i++ {
		nums[i] = 0
	}
}
