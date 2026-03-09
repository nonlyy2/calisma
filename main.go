package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{s1: "ab", s2: "eidbaooo", expected: true},  // "ba" — это перестановка "ab"
		{s1: "ab", s2: "eidboaoo", expected: false}, // буквы есть, но не подряд
		{s1: "adc", s2: "dcda", expected: true},     // "cda" — перестановка "adc"
		{s1: "hello", s2: "ooolleoooleh", expected: false},
		{s1: "a", s2: "ab", expected: true},
	}

	fmt.Println("🔄 Тестируем Permutation in String (#567)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.CheckInclusion(tc.s1, tc.s2)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (%s в %s)\n", i+1, tc.s1, tc.s2)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v для (%s в %s)\n", i+1, tc.expected, result, tc.s1, tc.s2)
		}
	}
}
