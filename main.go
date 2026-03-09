package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		paths    [][]string
		expected string
	}{
		{
			paths:    [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			expected: "Sao Paulo",
		},
		{
			paths:    [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			expected: "A",
		},
		{
			paths:    [][]string{{"A", "Z"}},
			expected: "Z",
		},
	}

	fmt.Println("✈️ Тестируем Destination City (#1436)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.DestCity(tc.paths)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %s\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %s, получили %s\n", i+1, tc.expected, result)
		}
	}
}
