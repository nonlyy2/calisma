package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Классический сдвиг вправо
	printTest([]int{3, 4, 5, 1, 2}, 1, 1)

	// ТЕСТ 2: Сдвиг ближе к середине
	printTest([]int{4, 5, 6, 7, 0, 1, 2}, 0, 2)

	// ТЕСТ 3: Массив изначально отсортирован (сдвига нет)
	printTest([]int{11, 13, 15, 17}, 11, 3)

	// ТЕСТ 4: Всего два элемента (граничный случай)
	printTest([]int{2, 1}, 1, 4)

	// ТЕСТ 5: Один элемент (граничный случай)
	printTest([]int{1}, 1, 5)
}

func printTest(nums []int, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)

	// Если функция в отдельном пакете, используй:
	// result := leetcode.FindMin(nums)
	result := leetcode.FindMin(nums)

	fmt.Printf("Input: %v | Result: %d (Expected: %d)\n", nums, result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
