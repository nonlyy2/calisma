package leetcode

// Задача №1: Two Sum
// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) [][]int {
	result := [][]int{}
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (num + nums[j]) == target {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
