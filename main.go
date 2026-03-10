package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	testCases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	fmt.Println("🚀 Тестируем Merge Sorted Array (#88)...")
	fmt.Println("---")

	for i, tc := range testCases {
		// Делаем копию nums1, так как функция меняет его in-place
		nums1Copy := make([]int, len(tc.nums1))
		copy(nums1Copy, tc.nums1)

		leetcode.Merge(nums1Copy, tc.m, tc.nums2, tc.n)

		if reflect.DeepEqual(nums1Copy, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден. Итог: %v\n", i+1, nums1Copy)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, nums1Copy)
		}
	}
}
