package leetcode

// Задача №424: Longest Repeating Character Replacement
// https://leetcode.com/problems/longest-repeating-character-replacement/

func CharacterReplacement(s string, k int) int {
	var counts [26]int
	maxFreq := 0
	res := 0
	left := 0

	for right := 0; right < len(s); right++ {
		counts[s[right]-'A']++
		maxFreq = max(maxFreq, counts[s[right]-'A'])
		if right-left-maxFreq >= k {
			counts[s[left]-'A']--
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}
