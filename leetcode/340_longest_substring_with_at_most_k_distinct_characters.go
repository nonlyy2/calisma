package leetcode

// Задача №340: Longest Substring with At Most K Distinct Characters
// https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/

func LengthOfLongestSubstringKDistinct(s string, k int) int {
	charMap := make(map[byte]int)
	maxLen := 0

	left := 0
	for right := range s {
		charMap[s[right]]++

		for len(charMap) > k {
			charMap[s[left]]--
			if charMap[s[left]] == 0 {
				delete(charMap, s[left])
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
