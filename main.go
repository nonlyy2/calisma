package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	tests := []struct {
		s        string
		t        string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"ab", "b", "b"},
		{"DONOTPANIC", "OOC", "ONOTPANIC"},
	}

	fmt.Println("🧪 Тестируем Minimum Window Substring...")

	for i, tc := range tests {
		result := leetcode.MinWindow(tc.s, tc.t)
		if result == tc.expected {
			fmt.Printf("✅ Тест %d: Пройден\n", i+1)
		} else {
			fmt.Printf("❌ Тест %d: ОШИБКА! Вход: s=\"%s\", t=\"%s\". Ожидали: \"%s\", Получили: \"%s\"\n",
				i+1, tc.s, tc.t, tc.expected, result)
		}
	}
}
