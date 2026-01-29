package leetcode

// Задача №128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	HashMap := make(map[int]bool)
	for _, val := range nums {
		HashMap[val] = false
	}

	for _, num := range nums {
		if _, exists := HashMap[num-1]; !exists {
			HashMap[num] = true
		}
	}

	max_streak := 0

	for val := range HashMap {
		if HashMap[val] == true {
			current_streak := 1
			current_num := val

			for {
				if _, exists := HashMap[current_num+1]; exists {
					current_streak++
					current_num++
				} else {
					break
				}
			}

			if current_streak > max_streak {
				max_streak = current_streak
			}
		}
	}

	return max_streak
}
