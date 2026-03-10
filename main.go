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
			nums:     []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			nums:     []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}

	fmt.Println("🔢 Тестируем Squares of a Sorted Array (#977)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.SortedSquares(tc.nums)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %v\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
