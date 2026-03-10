package leetcode

// Задача №2743: Count Substrings Without Repeating Character
// https://leetcode.com/problems/count-substrings-without-repeating-character/

func CountSubstrings(s string) int {
	charMap := [26]int{}
	count := 0

	left := 0
	for right := range s {
		charMap[s[right]-'a']++

		for charMap[s[right]-'a'] > 1 {
			charMap[s[left]-'a']--
			left++
		}

		count += right - left + 1
	}

	return count
}
