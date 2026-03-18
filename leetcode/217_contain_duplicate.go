package leetcode

// Задача №217: Contains Duplicate
// https://leetcode.com/problems/contains-duplicate/

func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, value := range nums {
		if _, exists := seen[value]; exists {
			return true
		} else {
			seen[value] = true
		}
	}
	return false
}
