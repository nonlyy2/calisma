package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	testCases := []struct {
		s        string
		p        string
		expected []int
	}{
		{
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6}, // "cba" на 0 индексе, "bac" на 6-м
		},
		{
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2}, // "ab", "ba", "ab" — окна могут пересекаться!
		},
		{
			s:        "af",
			p:        "be",
			expected: []int{}, // Анаграмм нет
		},
		{
			s:        "aaaaaaaaaa",
			p:        "aaaaaaaaaa",
			expected: []int{0},
		},
		{
			s:        "abc",
			p:        "abcd",
			expected: []int{}, // p длиннее чем s — сразу пустой результат
		},
	}

	fmt.Println("🔍 Тестируем Find All Anagrams (#438)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.FindAnagrams(tc.s, tc.p)

		// Используем reflect.DeepEqual для сравнения слайсов
		if (len(result) == 0 && len(tc.expected) == 0) || reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден (s=\"%s\", p=\"%s\")\n", i+1, tc.s, tc.p)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v (s=\"%s\", p=\"%s\")\n",
				i+1, tc.expected, result, tc.s, tc.p)
		}
	}
}
