package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Классический пример
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := 2

	fmt.Println("--- Тест 1 ---")
	fmt.Printf("Вход: nums = %v, k = %d\n", nums1, k1)

	// Вызываем твою функцию
	result1 := leetcode.TopKFrequent(nums1, k1)

	fmt.Printf("Результат: %v\n", result1) // Ожидаем [1, 2] (порядок не важен)

	fmt.Println() // Просто пустая строка для красоты

	// ТЕСТ 2: Один элемент
	nums2 := []int{1}
	k2 := 1

	fmt.Println("--- Тест 2 ---")
	fmt.Printf("Вход: nums = %v, k = %d\n", nums2, k2)

	result2 := leetcode.TopKFrequent(nums2, k2)

	fmt.Printf("Результат: %v\n", result2) // Ожидаем [1]

	fmt.Println()

	// ТЕСТ 3: Мой пример (чтобы проверить сортировку)
	// 10 встречается 3 раза, 20 встречается 2 раза, 30 - 1 раз.
	nums3 := []int{30, 20, 20, 10, 10, 10}
	k3 := 2

	fmt.Println("--- Тест 3 ---")
	fmt.Printf("Вход: nums = %v, k = %d\n", nums3, k3)

	result3 := leetcode.TopKFrequent(nums3, k3)

	fmt.Printf("Результат: %v\n", result3) // Ожидаем [10, 20]
}
