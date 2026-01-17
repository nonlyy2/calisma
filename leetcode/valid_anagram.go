package leetcode

// Задача №242: Valid Anagram
// https://leetcode.com/problems/valid-anagram/

func IsAnagram(s string, t string) bool {
	charMap := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		charMap[s[i]]++
	}

	for j := 0; j < len(s); j++ {
		charMap[t[j]]--
	}

	for _, count := range charMap {
		if count != 0 {
			return false
		}
	}

	return true
}
