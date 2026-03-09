package leetcode

// Задача №161: One Edit Distance
// https://leetcode.com/problems/one-edit-distance/

func IsOneEditDistance(s string, t string) bool {
	lenS, lenT := len(s), len(t)

	if lenS > lenT {
		return IsOneEditDistance(t, s)
	}

	if lenT-lenS > 1 {
		return false
	}

	for i := 0; i < lenS; i++ {
		if s[i] != t[i] {
			if lenS == lenT {
				return s[i+1:] == t[i+1:]
			} else {
				return s[i:] == t[i+1:]
			}
		}
	}

	return lenS+1 == lenT
}
