package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		s        string
		expected int
	}{
		{s: "abcd", expected: 10}, // a, b, c, d, ab, bc, cd, abc, bcd, abcd
		{s: "ooo", expected: 3},   // o, o, o
		{s: "abab", expected: 7},  // a, b, a, b, ab, ba, ab
	}

	fmt.Println("🔢 Тестируем Count Substrings (#2743)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.CountSubstrings(tc.s)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %d\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d\n", i+1, tc.expected, result)
		}
	}
}
