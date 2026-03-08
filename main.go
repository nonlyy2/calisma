package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 1, 0, 1},
			expected: 3, // Удаляем ноль, получаем три единицы
		},
		{
			nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5, // Удаляем второй ноль, склеиваем 1,1,1 и 1,1
		},
		{
			nums:     []int{1, 1, 1},
			expected: 2, // Обязаны удалить один элемент, останется 2
		},
		{
			nums:     []int{0, 0, 0},
			expected: 0, // Как ни крути, одни нули
		},
	}

	fmt.Println("🪟 Тестируем Longest Subarray of 1's (#1493)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.LongestSubarray(tc.nums)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (Длина: %d)\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d\n", i+1, tc.expected, result)
		}
	}
}
