package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	fmt.Println("🎲 Тестируем RandomizedSet (#380)...")
	fmt.Println("---")

	obj := leetcode.Constructor()

	// Тест 1: Вставка 1
	fmt.Printf("Insert(1): %v (ожидаем true)\n", obj.Insert(1))

	// Тест 2: Удаление 2 (которого нет)
	fmt.Printf("Remove(2): %v (ожидаем false)\n", obj.Remove(2))

	// Тест 3: Вставка 2
	fmt.Printf("Insert(2): %v (ожидаем true)\n", obj.Insert(2))

	// Тест 4: GetRandom (должен вернуть 1 или 2)
	fmt.Printf("GetRandom(): %d\n", obj.GetRandom())

	// Тест 5: Удаление 1
	fmt.Printf("Remove(1): %v (ожидаем true)\n", obj.Remove(1))

	// Тест 6: Проверка на наличие 2 после удаления 1
	fmt.Printf("Insert(2): %v (ожидаем false)\n", obj.Insert(2))

	// Тест 7: Финальный GetRandom (должен вернуть точно 2)
	fmt.Printf("GetRandom(): %d\n", obj.GetRandom())

	fmt.Println("---")
	fmt.Println("Если паники не было и булевы значения совпали — ты красава!")
}
