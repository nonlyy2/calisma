package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Классический пример
	// Лучшая пара: 8 (индекс 1) и 7 (индекс 8).
	// Ширина: 8-1 = 7. Высота: min(8, 7) = 7. Площадь: 49.
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	printTest(height1, 49, 1)

	// ТЕСТ 2: Просто две палки
	// Ширина 1, высота 1 -> Площадь 1
	height2 := []int{1, 1}
	printTest(height2, 1, 2)

	// ТЕСТ 3: Проверка логики сдвига
	// Сначала возьмет края (4 и 4, ширина 4) -> 16
	height3 := []int{4, 3, 2, 1, 4}
	printTest(height3, 16, 3)
}

func printTest(input []int, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)
	fmt.Printf("Input: %v\n", input)

	result := leetcode.MaxArea(input)
	fmt.Printf("Result: %d (Expected: %d)\n", result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
