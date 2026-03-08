package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		s        string
		expected bool
	}{
		{s: "aba", expected: true},                     // Уже палиндром
		{s: "abca", expected: true},                    // Удаляем 'c' (или 'b') -> "aba" или "aca"
		{s: "abc", expected: false},                    // Удаление одного символа не спасет
		{s: "eeccccbebaeeabebccceea", expected: false}, // Сложный тест на длинную строку
		{s: "cbbcc", expected: true},                   // Удаляем первую 'c' -> "bbcc" (нет), удаляем последнюю 'c' -> "cbbc" (да!)
	}

	fmt.Println("🔍 Тестируем Valid Palindrome II (#680)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.ValidPalindrome(tc.s)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (\"%s\")\n", i+1, tc.s)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Строка: \"%s\". Ожидали %v, получили %v\n", i+1, tc.s, tc.expected, result)
		}
	}
}
