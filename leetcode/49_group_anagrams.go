package leetcode

// Задача №49: Group Anagrams
// https://leetcode.com/problems/group-anagrams/

func GroupAnagrams(strs []string) [][]string {
	HashMap := make(map[[26]int][]string)

	for _, word := range strs {
		var count [26]int

		for _, ch := range word {
			count[ch-'a']++
		}

		HashMap[count] = append(HashMap[count], word)
	}

	var result [][]string

	for _, group := range HashMap {
		result = append(result, group)
	}

	return result
}
