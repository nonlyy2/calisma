package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		seats    []int
		expected int
	}{
		{
			// Классический случай: садимся посередине (индекс 2)
			seats:    []int{1, 0, 0, 0, 1, 0, 1},
			expected: 2,
		},
		{
			// Краевой случай 1: нули в самом конце. Садимся на последний стул (индекс 3)
			seats:    []int{1, 0, 0, 0},
			expected: 3,
		},
		{
			// Краевой случай 2: нули в самом начале. Садимся на первый стул (индекс 0)
			seats:    []int{0, 0, 1},
			expected: 2,
		},
		{
			// Минимально возможное окно
			seats:    []int{1, 0, 1},
			expected: 1,
		},
		{
			// Длинная пустая дистанция с краю
			seats:    []int{0, 0, 0, 0, 1, 0, 1},
			expected: 4,
		},
	}

	fmt.Println("🍿 Тестируем Maximize Distance to Closest Person (#849)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.MaxDistToClosest(tc.seats)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (Макс. дистанция: %d)\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Места: %v. Ожидали %d, получили %d\n", i+1, tc.seats, tc.expected, result)
		}
	}
}
