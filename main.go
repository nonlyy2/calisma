package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		s        string
		k        int
		expected int
	}{
		{s: "eceba", k: 2, expected: 3},  // "ece"
		{s: "aa", k: 1, expected: 2},     // "aa"
		{s: "a", k: 0, expected: 0},      // k = 0, ни одна подстрока не подходит
		{s: "WORLD", k: 10, expected: 5}, // k больше длины строки
		{s: "abaccc", k: 2, expected: 4}, // "accc"
		{s: "", k: 3, expected: 0},       // пустая строка
	}

	fmt.Println("📏 Тестируем Longest Substring with At Most K Distinct Characters (#340)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.LengthOfLongestSubstringKDistinct(tc.s, tc.k)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %d\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d для строки %q (k=%d)\n", i+1, tc.expected, result, tc.s, tc.k)
		}
	}
}
