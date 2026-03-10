package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	fmt.Println("📚 Тестируем Min Stack (#155)...")
	fmt.Println("-------------------------------------------")

	obj := leetcode.ConstructorStack()

	// Сценарий 1: Базовые операции
	fmt.Println("Шаг 1: Push(-2), Push(0), Push(-3)")
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	check("GetMin() после пуша -3", obj.GetMin(), -3)

	fmt.Println("Шаг 2: Pop()")
	obj.Pop()

	check("Top() после Pop()", obj.Top(), 0)
	check("GetMin() после Pop()", obj.GetMin(), -2)

	// Сценарий 2: Еще пуши и проверки
	fmt.Println("Шаг 3: Push(-5), Push(2)")
	obj.Push(-5)
	obj.Push(2)

	check("GetMin() после пуша -5", obj.GetMin(), -5)

	fmt.Println("Шаг 4: Pop(), Pop()")
	obj.Pop()
	obj.Pop()

	check("GetMin() должен вернуться к -2", obj.GetMin(), -2)

	fmt.Println("-------------------------------------------")
	fmt.Println("Все тесты завершены.")
}

// вспомогательная функция для "птичек"
func check(name string, got, want int) {
	if got == want {
		fmt.Printf("✅ %s: Получено %d\n", name, got)
	} else {
		fmt.Printf("❌ %s: ОШИБКА! Ожидали %d, получили %d\n", name, want, got)
	}
}
