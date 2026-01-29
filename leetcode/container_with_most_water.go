package leetcode

// Задача №11: Container With Most Water
// https://leetcode.com/problems/container-with-most-water/

func MaxArea(height []int) int {
	l := 0
	r := len(height) - 1

	max_area := 0
	for l < r {
		area := (r - l) * min(height[l], height[r])

		if area > max_area {
			max_area = area
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max_area
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
