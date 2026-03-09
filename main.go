package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "ab", t: "acb", expected: true},    // Вставка
		{s: "cab", t: "ad", expected: false},   // Слишком много правок
		{s: "1203", t: "1213", expected: true}, // Замена
		{s: "", t: "A", expected: true},        // Вставка в пустую
		{s: "abc", t: "abc", expected: false},  // Одинаковые
	}

	fmt.Println("🔍 Тестируем One Edit Distance (#161)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.IsOneEditDistance(tc.s, tc.t)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (%s, %s)\n", i+1, tc.s, tc.t)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v для (%s, %s)\n", i+1, tc.expected, result, tc.s, tc.t)
		}
	}
}
