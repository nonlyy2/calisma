package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	fmt.Println("🔄 Тестируем Longest Palindromic Substring (#5)...")
	fmt.Println("-------------------------------------------")

	testCases := []struct {
		s        string
		expected []string // Несколько вариантов, если длина совпадает
	}{
		{s: "babad", expected: []string{"bab", "aba"}},
		{s: "cbbd", expected: []string{"bb"}},
		{s: "a", expected: []string{"a"}},
		{s: "ac", expected: []string{"a", "c"}},
		{s: "racecar", expected: []string{"racecar"}},
	}

	for i, tc := range testCases {
		result := leetcode.LongestPalindrome(tc.s)

		match := false
		for _, exp := range tc.expected {
			if result == exp {
				match = true
				break
			}
		}

		if match {
			fmt.Printf("✅ Тест %d: '%s' -> '%s'\n", i+1, tc.s, result)
		} else {
			fmt.Printf("❌ Тест %d: '%s' -> ОШИБКА! Получили '%s', ждали один из %v\n", i+1, tc.s, result, tc.expected)
		}
	}
}
