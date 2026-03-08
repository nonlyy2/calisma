package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		points   [][]int
		expected bool
	}{
		{
			// Пример 1: Симметрично относительно x = 0
			points:   [][]int{{1, 1}, {-1, 1}},
			expected: true,
		},
		{
			// Пример 2: Несимметрично (разная высота y)
			points:   [][]int{{1, 1}, {-1, 2}},
			expected: false,
		},
		{
			// Пример 3: Три точки, одна на оси симметрии
			points:   [][]int{{1, 1}, {-1, 1}, {0, 1}},
			expected: true,
		},
		{
			// Пример 4: Точки с одинаковыми координатами (дубликаты)
			points:   [][]int{{1, 1}, {1, 1}, {-1, 1}},
			expected: true,
		},
		{
			// Пример 5: Сложный случай
			points:   [][]int{{0, 0}, {1, 0}, {3, 0}},
			expected: false,
		},
	}

	fmt.Println("🪞 Тестируем Line Reflection (#356)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.IsReflected(tc.points)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден\n", i+1)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
