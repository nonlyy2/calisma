package leetcode

// Задача №567: Permutation in String
// https://leetcode.com/problems/permutation-in-string/

func CheckInclusion(s1 string, s2 string) bool {
	targetSet := [26]int{}
	charSet := [26]int{}

	for i := 0; i < len(s1); i++ {
		targetSet[s1[i]-'a']++
	}

	left := 0

	for right := 0; right < len(s2); right++ {
		charSet[s2[right]-'a']++

		if right-left+1 > len(s1) {
			charSet[s2[left]-'a']--
			left++
		}

		if charSet == targetSet {
			return true
		}

	}

	return false
}
