package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь к пакету
	"reflect"
)

func main() {
	testCases := []struct {
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 2, 3, 4},
		},
		{
			arr:      []int{1, 1, 1, 10, 10, 10},
			k:        1,
			x:        9,
			expected: []int{10},
		},
		{
			arr:      []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
			k:        3,
			x:        5,
			expected: []int{3, 3, 4},
		},
	}

	fmt.Println("🔍 Тестируем Find K Closest Elements (#658)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.FindClosestElements(tc.arr, tc.k, tc.x)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден (k=%d, x=%d)\n", i+1, tc.k, tc.x)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
