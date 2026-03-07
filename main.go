package main

import (
	"fmt"
	"neetcode/leetcode" // Твой путь к пакету
)

func main() {
	testCases := []struct {
		grid     [][]byte
		expected int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			grid: [][]byte{
				{'1', '0', '1', '0', '1'},
			},
			expected: 3,
		},
		{
			grid:     [][]byte{},
			expected: 0,
		},
	}

	fmt.Println("🏝️ Тестируем Number of Islands (#200)...")
	fmt.Println("---")

	for i, tc := range testCases {
		// Делаем копию грида, так как алгоритм может его изменять (мутировать)
		gridCopy := make([][]byte, len(tc.grid))
		for j := range tc.grid {
			gridCopy[j] = make([]byte, len(tc.grid[j]))
			copy(gridCopy[j], tc.grid[j])
		}

		result := leetcode.NumIslands(gridCopy)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (Ожидали %d, получили %d)\n", i+1, tc.expected, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d\n", i+1, tc.expected, result)
		}
	}
}
