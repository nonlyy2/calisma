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
		{nums: []int{1, 4, 3, 3, 2}, expected: 2}, // [1,4] или [3,2]
		{nums: []int{3, 3, 3, 3}, expected: 1},    // Нет строгого возрастания/убывания
		{nums: []int{3, 2, 1}, expected: 3},       // Строго убывает
		{nums: []int{1, 2, 3, 4, 5}, expected: 5}, // Строго возрастает
		{nums: []int{1, 5, 2, 7, 3}, expected: 2},
	}

	fmt.Println("🔍 Тестируем Longest Subarray (#3105)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.LongestMonotonicSubarray(tc.nums)
		if result == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (nums=%v)\n", i+1, tc.nums)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %d\n", i+1, tc.expected, result)
		}
	}
}
