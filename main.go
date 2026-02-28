package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Таргет в правой (сдвинутой) части
	printTest([]int{4, 5, 6, 7, 0, 1, 2}, 0, 4, 1)

	// ТЕСТ 2: Таргет не существует
	printTest([]int{4, 5, 6, 7, 0, 1, 2}, 3, -1, 2)

	// ТЕСТ 3: Обычный массив (без сдвига), таргет в левой части
	printTest([]int{1, 2, 3, 4, 5}, 2, 1, 3)

	// ТЕСТ 4: Один элемент (нашли)
	printTest([]int{1}, 1, 0, 4)

	// ТЕСТ 5: Один элемент (не нашли)
	printTest([]int{1}, 0, -1, 5)

	// ТЕСТ 6: Два элемента, таргет во второй половине
	printTest([]int{3, 1}, 1, 1, 6)

	// ТЕСТ 7: Таргет на границе разлома
	printTest([]int{5, 1, 3}, 5, 0, 7)
}

func printTest(nums []int, target, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)

	// Если перенесешь функцию в отдельный файл/пакет:
	result := leetcode.Search(nums, target)

	fmt.Printf("Input: %v, target: %d | Result: %d (Expected: %d)\n", nums, target, result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
