package leetcode

// Задача №167: Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func TwoSumII(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		sum := numbers[l] + numbers[r]

		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else if sum > target {
			r--
		}
	}

	return nil
}
