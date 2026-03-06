package leetcode

// Задача №205: Isomorphic Strings
// https://leetcode.com/problems/isomorphic-strings/

func IsIsomorphic(s string, t string) bool {
	hashMap := make(map[byte]byte)
	hashMapReversed := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		val1, exists1 := hashMap[s[i]]
		val2, exists2 := hashMapReversed[t[i]]

		if exists1 || exists2 {
			if t[i] != val1 || s[i] != val2 {
				return false
			}
		} else {
			hashMap[s[i]] = t[i]
			hashMapReversed[t[i]] = s[i]
		}

	}

	return true
}
