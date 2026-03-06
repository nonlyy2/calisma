package leetcode

// Задача №228: Summary Ranges
// https://leetcode.com/problems/summary-ranges/

import (
	"fmt"
	"strconv"
)

func SummaryRanges(nums []int) []string {
	res := []string{}
	start := 0
	count := 0

	for i, _ := range nums {
		if i != len(nums)-1 && nums[i]+1 == nums[i+1] {
			count++
		} else {
			if count == 0 {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[start+count]))
			}

			start = i + 1
			count = 0
		}
	}

	return res
}
