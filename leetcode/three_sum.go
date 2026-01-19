package leetcode

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	var l, r, sum int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l = i + 1
		r = len(nums) - 1

		for l < r {
			sum = nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
