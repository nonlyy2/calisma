package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	fmt.Println("🔄 Тестируем Linked List Cycle (#141)...")
	fmt.Println("-------------------------------------------")

	// Тест 1: Список с циклом 3 -> 2 -> 0 -> -4 -> (назад к 2)
	node1 := &leetcode.ListNode{Val: 3}
	node2 := &leetcode.ListNode{Val: 2}
	node3 := &leetcode.ListNode{Val: 0}
	node4 := &leetcode.ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // Создаем цикл

	fmt.Printf("Тест 1 (с циклом): Получено %v, ожидаем true ✅\n", leetcode.HasCycle(node1))

	// Тест 2: Список без цикла 1 -> 2
	nodeA := &leetcode.ListNode{Val: 1}
	nodeB := &leetcode.ListNode{Val: 2}
	nodeA.Next = nodeB

	fmt.Printf("Тест 2 (без цикла): Получено %v, ожидаем false ✅\n", leetcode.HasCycle(nodeA))

	fmt.Println("-------------------------------------------")
}
