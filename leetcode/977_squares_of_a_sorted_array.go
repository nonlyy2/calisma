package leetcode

// Задача №977: Squares of a Sorted Array
// https://leetcode.com/problems/squares-of-a-sorted-array/

func SortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	currIndex := length - 1

	left, right := 0, length-1
	for left <= right {
		v1, v2 := nums[left], nums[right]

		if v1*v1 <= v2*v2 {
			res[currIndex] = v2 * v2
			right--
		} else {
			res[currIndex] = v1 * v1
			left++
		}

		currIndex--
	}

	return res
}
