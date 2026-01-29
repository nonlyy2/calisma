package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: База
	printTest("()", true, 1)

	// ТЕСТ 2: Разные типы
	printTest("()[]{}", true, 2)

	// ТЕСТ 3: Не та пара
	printTest("(]", false, 3)

	// ТЕСТ 4: Правильная вложенность
	printTest("{[]}", true, 4)

	// ТЕСТ 5: Лишняя открывающая (стек не пуст в конце)
	printTest("[", false, 5)

	// ТЕСТ 6: Лишняя закрывающая (стек пуст, а мы хотим попнуть)
	printTest("]", false, 6)
}

func printTest(input string, expected bool, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)
	fmt.Printf("Input: \"%s\"\n", input)

	result := leetcode.IsValid(input)
	fmt.Printf("Result: %v (Expected: %v)\n", result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
