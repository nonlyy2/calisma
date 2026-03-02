package leetcode

// Задача №76: Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/

func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var targetCount [128]int
	var windowCount [128]int
	need := 0

	for i := 0; i < len(t); i++ {
		if targetCount[t[i]] == 0 {
			need++
		}
		targetCount[t[i]]++
	}

	left := 0
	have := 0
	minLen := len(s) + 1
	start := 0

	for right := 0; right < len(s); right++ {
		windowCount[s[right]]++

		if windowCount[s[right]] == targetCount[s[right]] {
			have++
		}

		for have == need {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			if targetCount[s[left]] > 0 && windowCount[s[left]] == targetCount[s[left]] {
				have--
			}

			windowCount[s[left]]--
			left++
		}

	}

	if minLen == len(s)+1 {
		return ""
	}

	res := s[start : start+minLen]

	return res
}
