package main

import (
	"fmt"
	"neetcode/leetcode"
)

// sliceToList конвертирует обычный слайс в связный список
func sliceToList(nums []int) *leetcode.ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &leetcode.ListNode{}
	curr := dummy
	for _, n := range nums {
		curr.Next = &leetcode.ListNode{Val: n}
		curr = curr.Next
	}
	return dummy.Next
}

// listToSlice конвертирует список обратно в слайс (удобно для сравнения в тестах)
func listToSlice(head *leetcode.ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// printList печатает список в консоль в удобном виде: 1 -> 2 -> 3
func printList(head *leetcode.ListNode) {
	if head == nil {
		fmt.Println("Empty List (nil)")
		return
	}
	curr := head
	for curr != nil {
		fmt.Printf("%d", curr.Val)
		if curr.Next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println("🔗 Тестируем Remove Nth Node From End (#19)...")
	fmt.Println("-------------------------------------------")

	testCases := []struct {
		nums     []int
		n        int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{1, 2, 3, 5},
		},
		{
			nums:     []int{1},
			n:        1,
			expected: []int{},
		},
		{
			nums:     []int{1, 2},
			n:        1,
			expected: []int{1},
		},
		{
			nums:     []int{1, 2},
			n:        2,
			expected: []int{2},
		},
	}

	for i, tc := range testCases {
		head := sliceToList(tc.nums)
		fmt.Printf("Тест %d: Вход: %v, n: %d\n", i+1, tc.nums, tc.n)

		result := leetcode.RemoveNthFromEnd(head, tc.n)

		fmt.Print("       Результат: ")
		printList(result)

		// Простая визуальная проверка (можно добавить reflect.DeepEqual для строгости)
		fmt.Println("-------------------------------------------")
	}
}
