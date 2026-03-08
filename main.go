package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	testCases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2, // [1,1] и [1,1]
		},
		{
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2, // [1,2] и [3]
		},
		{
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 3, // [1,-1], [0], [1,-1,0]
		},
		{
			nums:     []int{-1, -1, 1},
			k:        0,
			expected: 1, // [-1, 1]
		},
	}

	fmt.Println("🔢 Тестируем Subarray Sum Equals K (#560)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.SubarraySum(tc.nums, tc.k)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %d\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d\n", i+1, tc.expected, result)
		}
	}
}
