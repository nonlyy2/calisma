package leetcode

// Задача №5: Longest Palindromic Substring
// https://leetcode.com/problems/longest-palindromic-substring/

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandFromCenter(s, i, i)
		len2 := expandFromCenter(s, i, i+1)

		maxLen := max(len1, len2)

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}
