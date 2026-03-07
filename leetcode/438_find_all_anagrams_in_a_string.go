package leetcode

// Задача №438: Find All Anagrams in a String
// https://leetcode.com/problems/find-all-anagrams-in-a-string/

func FindAnagrams(s string, p string) []int {
	var sCount [26]int
	var pCount [26]int
	var res []int

	for i := range p {
		pCount[p[i]-'a']++
	}

	left := 0
	for right := 0; right < len(s); right++ {
		sCount[s[right]-'a']++

		if right-left+1 == len(p) {
			if sCount == pCount {
				res = append(res, left)
			}

			sCount[s[left]-'a']--
			left++
		}
	}

	return res
}
