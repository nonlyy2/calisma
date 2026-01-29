package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Классика
	// Цепочка: 1, 2, 3, 4
	nums1 := []int{100, 4, 200, 1, 3, 2}
	printTest(nums1, 4, 1)

	// ТЕСТ 2: Длинная змейка
	// Цепочка: 0, 1, 2, 3, 4, 5, 6, 7, 8
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	printTest(nums2, 9, 2)

	// ТЕСТ 3: Пустой массив
	nums3 := []int{}
	printTest(nums3, 0, 3)

	// ТЕСТ 4: Только одно число (или дубликаты одного)
	nums4 := []int{1, 1, 1, 1} // Длина 1
	printTest(nums4, 1, 4)
}

func printTest(input []int, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)
	fmt.Printf("Input: %v\n", input)

	result := leetcode.LongestConsecutive(input)
	fmt.Printf("Result: %d (Expected: %d)\n", result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
