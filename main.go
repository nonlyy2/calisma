package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Классический пример из описания
	// 1*2*3*4 = 24. Для каждого элемента убираем его самого.
	nums1 := []int{1, 2, 3, 4}

	fmt.Println("--- Тест 1 (Обычный) ---")
	fmt.Printf("Вход: nums = %v\n", nums1)

	// Вызываем твою функцию
	result1 := leetcode.ProductExceptSelf(nums1)

	fmt.Printf("Результат: %v\n", result1)
	// Ожидаем: [24, 12, 8, 6]

	fmt.Println()

	// ТЕСТ 2: Массив с одним нулем
	// Это самый хитрый случай. Там, где 0, будет число. В остальных местах - 0.
	nums2 := []int{-1, 1, 0, -3, 3}

	fmt.Println("--- Тест 2 (Один ноль) ---")
	fmt.Printf("Вход: nums = %v\n", nums2)

	result2 := leetcode.ProductExceptSelf(nums2)

	fmt.Printf("Результат: %v\n", result2)
	// Ожидаем: [0, 0, 9, 0, 0]
	// Пояснение: (-1)*1*(-3)*3 = 9. Остальные умножаются на 0.

	fmt.Println()

	// ТЕСТ 3: Массив с двумя нулями
	// Если нулей больше одного, результат ВСЕГДА будет из всех нулей.
	nums3 := []int{0, 4, 0}

	fmt.Println("--- Тест 3 (Два нуля) ---")
	fmt.Printf("Вход: nums = %v\n", nums3)

	result3 := leetcode.ProductExceptSelf(nums3)

	fmt.Printf("Результат: %v\n", result3)
	// Ожидаем: [0, 0, 0]
}
