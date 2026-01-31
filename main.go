package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Классика (ответ 3 -> "abc")
	printTest("abcabcbb", 3, 1)

	// ТЕСТ 2: Все одинаковые (ответ 1 -> "b")
	printTest("bbbbb", 1, 2)

	// ТЕСТ 3: Ответ в конце и с перехлестом (ответ 3 -> "wke")
	// Подстрока обязательно должна быть слитной, "pwke" - это подпоследовательность, не считается.
	printTest("pwwkew", 3, 3)

	// ТЕСТ 4: Пустая строка
	printTest("", 0, 4)

	// ТЕСТ 5: Пробелы тоже символы
	printTest(" ", 1, 5)

	// ТЕСТ 6:
	printTest("qrsvbspk", 5, 6)
}

func printTest(input string, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)
	result := leetcode.LengthOfLongestSubstring(input)
	fmt.Printf("Input: \"%s\" | Result: %d (Expected: %d)\n", input, result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
