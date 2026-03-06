package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь к пакету
	"reflect"
)

func main() {
	testCases := []struct {
		nums     []int
		expected []string
	}{
		{
			nums:     []int{0, 1, 2, 4, 5, 7},
			expected: []string{"0->2", "4->5", "7"},
		},
		{
			nums:     []int{0, 2, 3, 4, 6, 8, 9},
			expected: []string{"0", "2->4", "6", "8->9"},
		},
		{
			nums:     []int{},
			expected: []string{},
		},
		{
			nums:     []int{-1},
			expected: []string{"-1"},
		},
		{
			nums:     []int{-5, -4, -3, 0, 1, 5},
			expected: []string{"-5->-3", "0->1", "5"},
		},
	}

	fmt.Println("🔍 Тестируем Summary Ranges (#228)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.SummaryRanges(tc.nums)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден (nums=%v)\n", i+1, tc.nums)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
