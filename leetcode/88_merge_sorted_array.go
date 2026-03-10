package leetcode

// Задача №88: Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	currPos := m + n - 1

	for j >= 0 && i >= 0 {
		if nums1[i] > nums2[j] {
			nums1[currPos] = nums1[i]
			i--
		} else {
			nums1[currPos] = nums2[j]
			j--
		}

		currPos--
	}

	for j >= 0 {
		nums1[currPos] = nums2[j]
		j--
		currPos--
	}
}
