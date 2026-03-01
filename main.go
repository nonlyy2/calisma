package main

import (
	"fmt"
	"neetcode/additional" // Проверь свой импорт, если нужно
)

func main() {
	// ТЕСТ 1: Обычный случай со всеми вариантами
	printTest("AAAABBBCCXYZ", "A4B3C2XYZ", 1)

	// ТЕСТ 2: Одиночные символы в начале и середине
	printTest("ABBBCHHH", "AB3CH3", 2)

	// ТЕСТ 3: Тот самый "потерянный хвост" (одиночный символ в конце)
	printTest("AAAB", "A3B", 3)

	// ТЕСТ 4: Пустая строка
	printTest("", "", 4)

	// ТЕСТ 5: Строка из одного символа
	printTest("A", "A", 5)

	// ТЕСТ 6: Все символы уникальные (худший случай для RLE)
	printTest("ABCDE", "ABCDE", 6)

	// ТЕСТ 7: Длинная последовательность (двузначное число > 9)
	printTest("AAAAAAAAAAAAA", "A13", 7)
}

func printTest(s, expected string, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)

	result := additional.RLE(s)

	fmt.Printf("Input: \"%s\"\nResult: \"%s\" (Expected: \"%s\")\n", s, result, expected)

	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
