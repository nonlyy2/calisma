package leetcode

// Задача №658: Find K Closest Elements
// https://leetcode.com/problems/find-k-closest-elements/

func FindClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k

	for right > left {
		mid := (left + right) / 2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
