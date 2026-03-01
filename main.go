package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Заменяем 2 символа
	// ABAB, k=2 -> BBBB (4)
	printTest("ABAB", 2, 4, 1)

	// ТЕСТ 2: Заменяем 1 символ
	// AABABBA, k=1 -> AAAABBA (4)
	printTest("AABABBA", 1, 4, 2)

	// ТЕСТ 3: k=0 (замены запрещены)
	// ABCDE -> 1
	printTest("ABCDE", 0, 1, 3)

	// ТЕСТ 4: Одна буква
	printTest("A", 2, 1, 4)

	// ТЕСТ 5: Строка из одинаковых букв
	printTest("AAAA", 2, 4, 5)

	// ТЕСТ 6: Замены позволяют объединить две группы
	// AAABBAAA, k=2 -> AAAAAAAA (8)
	printTest("AAABBAAA", 2, 8, 6)
}

func printTest(s string, k, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)

	result := leetcode.CharacterReplacement(s, k)

	fmt.Printf("Input: s=\"%s\", k=%d | Result: %d (Expected: %d)\n", s, k, result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
