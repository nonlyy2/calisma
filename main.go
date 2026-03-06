package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь к пакету
)

func main() {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{s: "egg", t: "add", expected: true},
		{s: "foo", t: "bar", expected: false},
		{s: "paper", t: "title", expected: true},
		{s: "badc", t: "baba", expected: false}, // Ловушка на обратный маппинг
		{s: "ab", t: "aa", expected: false},
	}

	fmt.Println("🔍 Тестируем Isomorphic Strings (#205)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.IsIsomorphic(tc.s, tc.t)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (\"%s\" -> \"%s\")\n", i+1, tc.s, tc.t)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v (\"%s\" -> \"%s\")\n",
				i+1, tc.expected, result, tc.s, tc.t)
		}
	}
}
