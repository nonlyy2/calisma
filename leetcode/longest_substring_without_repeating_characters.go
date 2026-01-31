package leetcode

// Задача №3: Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func LengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]bool)
	l := 0
	maxLen := 0

	for r := 0; r < len(s); r++ {
		for charMap[s[r]] {
			delete(charMap, s[l])
			l++
		}

		charMap[s[r]] = true
		if r-l+1 > maxLen {
			maxLen = len(charMap)
		}
	}

	return maxLen
}
