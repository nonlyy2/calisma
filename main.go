package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		grid     [][]int
		expected int
	}{
		{
			grid: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 1, 0, 1},
				{0, 1, 0, 0, 1},
				{0, 0, 0, 1, 1},
			},
			expected: 4, // Остров в середине (3 клетки) и остров в углу (4 клетки)
		},
		{
			grid: [][]int{
				{0, 0, 0, 0, 0},
			},
			expected: 0,
		},
		{
			grid: [][]int{
				{1, 1, 0, 1, 1},
				{1, 1, 0, 1, 1},
			},
			expected: 4,
		},
	}

	fmt.Println("🌊 Тестируем Max Area of Island (#695)...")
	fmt.Println("---")

	for i, tc := range testCases {
		// Копируем, чтобы не портить оригинал для других тестов
		gridCopy := make([][]int, len(tc.grid))
		for j := range tc.grid {
			gridCopy[j] = make([]int, len(tc.grid[j]))
			copy(gridCopy[j], tc.grid[j])
		}

		result := leetcode.MaxAreaOfIsland(gridCopy)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ (Max Area: %d)\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d\n", i+1, tc.expected, result)
		}
	}
}
