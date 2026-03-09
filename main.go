package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			nums:     []int{0},
			expected: []int{0},
		},
		{
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
	}

	fmt.Println("🚀 Тестируем Move Zeroes (#283)...")
	fmt.Println("---")

	for i, tc := range testCases {
		// Копируем слайс, чтобы не менять оригинал в тест-кейсе
		input := make([]int, len(tc.nums))
		copy(input, tc.nums)

		leetcode.MoveZeroes(input)

		if reflect.DeepEqual(input, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %v\n", i+1, input)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, input)
		}
	}
}
