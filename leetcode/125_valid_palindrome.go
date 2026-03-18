package leetcode

// Задача №125: Valid Palindrome
// https://leetcode.com/problems/valid-palindrome/

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	runes := []rune(s)
	l, r := 0, len(runes)-1

	for l < r {
		if !isAlphanumeric(runes[l]) {
			l++
			continue
		}
		if !isAlphanumeric(runes[r]) {
			r--
			continue
		}
		if unicode.ToLower(runes[r]) != unicode.ToLower(runes[l]) {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
