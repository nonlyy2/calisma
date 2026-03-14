package main

import (
	"fmt"
	"neetcode/leetcode"
)

// Вспомогательная функция для печати
func printList(head *leetcode.ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println("🔗 Тестируем Reorder List (#143)...")

	// Создаем 1 -> 2 -> 3 -> 4 -> 5
	n5 := &leetcode.ListNode{Val: 5}
	n4 := &leetcode.ListNode{Val: 4, Next: n5}
	n3 := &leetcode.ListNode{Val: 3, Next: n4}
	n2 := &leetcode.ListNode{Val: 2, Next: n3}
	n1 := &leetcode.ListNode{Val: 1, Next: n2}

	leetcode.ReorderList(n1)

	fmt.Print("Результат: ")
	printList(n1)
	fmt.Println("Ожидаем:   1 5 2 4 3")
}
