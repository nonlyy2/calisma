package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь
)

func main() {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "abc", t: "ahbgdc", expected: true},
		{s: "axc", t: "ahbgdc", expected: false},
		{s: "", t: "ahbgdc", expected: true}, // Пустая строка - всегда подпоследовательность
		{s: "abc", t: "", expected: false},
		{s: "aaaaaa", t: "bbaaaa", expected: false},
		{s: "aec", t: "abcde", expected: false}, // Порядок важен!
	}

	fmt.Println("🔍 Тестируем Is Subsequence (#392)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.IsSubsequence(tc.s, tc.t)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (s=\"%s\", t=\"%s\")\n", i+1, tc.s, tc.t)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
