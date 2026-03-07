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
		{s: "3+2*2", expected: 7},
		{s: " 3/2 ", expected: 1},
		{s: " 3+5 / 2 ", expected: 5},
		{s: "10-2*3+4", expected: 8},
		{s: "42", expected: 42},
		{s: "1-1+1", expected: 1},
	}

	fmt.Println("🧮 Тестируем Basic Calculator II (#227)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.Calculate(tc.s)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (\"%s\" = %d)\n", i+1, tc.s, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! (\"%s\") Ожидали %d, получили %d\n", i+1, tc.s, tc.expected, result)
		}
	}
}
