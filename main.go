package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	testCases := []struct {
		A        []int
		B        []int
		expected []int
	}{
		{
			A:        []int{1, 3, 2, 4},
			B:        []int{3, 1, 2, 4},
			expected: []int{0, 2, 3, 4},
		},
		{
			A:        []int{2, 3, 1},
			B:        []int{3, 1, 2},
			expected: []int{0, 1, 3},
		},
	}

	fmt.Println("📊 Тестируем Find the Prefix Common Array (#2657)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.FindThePrefixCommonArray(tc.A, tc.B)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %v\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
