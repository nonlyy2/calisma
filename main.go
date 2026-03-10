package main

import (
	"fmt"
	"neetcode/leetcode"
)

// Вспомогательная функция для создания списка из слайса
func sliceToList(nums []int) *leetcode.ListNode {
	dummy := &leetcode.ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &leetcode.ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// Вспомогательная функция для вывода списка
func printList(l *leetcode.ListNode) {
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println("➕ Тестируем Add Two Numbers (#2)...")
	fmt.Println("---")

	l1 := sliceToList([]int{2, 4, 3})
	l2 := sliceToList([]int{5, 6, 4})
	// Ожидаем: 7 0 8
	result := leetcode.AddTwoNumbers(l1, l2)
	fmt.Print("Тест 1: ")
	printList(result)

	l3 := sliceToList([]int{9, 9, 9})
	l4 := sliceToList([]int{1})
	// Ожидаем: 0 0 0 1
	result2 := leetcode.AddTwoNumbers(l3, l4)
	fmt.Print("Тест 2: ")
	printList(result2)
}
