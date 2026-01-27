package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	// ТЕСТ 1: Обычные слова
	input1 := []string{"neet", "code", "love", "you"}
	fmt.Println("--- Тест 1 (Базовый) ---")
	runTest(input1)

	// ТЕСТ 2: Символы, которые ломают простые сплиты
	// Тут есть и пробелы, и решетки, и цифры
	input2 := []string{"we", "say", ":", "yes", "!@#$%^&*()"}
	fmt.Println("\n--- Тест 2 (Спецсимволы) ---")
	runTest(input2)

	// ТЕСТ 3: Пустые строки и пустой список
	input3 := []string{""} // Список с одной пустой строкой
	fmt.Println("\n--- Тест 3 (Пустая строка) ---")
	runTest(input3)

	input4 := []string{} // Вообще пустой список
	fmt.Println("\n--- Тест 4 (Пустой список) ---")
	runTest(input4)
}

func runTest(input []string) {
	fmt.Printf("Input:  %v\n", input)

	encoded := leetcode.Encode(input)
	fmt.Printf("Encoded: %s\n", encoded)

	decoded := leetcode.Decode(encoded)
	fmt.Printf("Decoded: %v\n", decoded)

	if reflect.DeepEqual(input, decoded) {
		fmt.Println("✅ SUCCESS")
	} else {
		fmt.Println("❌ FAILED")
	}
}
