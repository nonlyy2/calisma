package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		chars    []string // используем string для удобства инициализации
		expected int
		result   []string
	}{
		{
			chars:    []string{"a", "a", "b", "b", "c", "c", "c"},
			expected: 6, // "a2b2c3"
			result:   []string{"a", "2", "b", "2", "c", "3"},
		},
		{
			chars:    []string{"a"},
			expected: 1, // "a"
			result:   []string{"a"},
		},
		{
			chars:    []string{"a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"},
			expected: 4, // "ab12"
			result:   []string{"a", "b", "1", "2"},
		},
	}

	fmt.Println("🗜️ Тестируем String Compression (#443)...")
	fmt.Println("---")

	for i, tc := range testCases {
		// Конвертируем []string в []byte для функции
		byteChars := make([]byte, len(tc.chars))
		for j, s := range tc.chars {
			byteChars[j] = s[0]
		}

		newLen := leetcode.Compress(byteChars)

		if newLen == tc.expected {
			fmt.Printf("Тест %d: ✅ Длина совпала (%d)\n", i+1, newLen)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка длины! Ожидали %d, получили %d\n", i+1, tc.expected, newLen)
		}
	}
}
