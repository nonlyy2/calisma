package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

// sliceToLinkedList преобразует []int в связный список
func sliceToLinkedList(nums []int) *leetcode.ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &leetcode.ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &leetcode.ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

// linkedListToSlice преобразует список обратно в []int для проверки
func linkedListToSlice(head *leetcode.ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func main() {
	fmt.Println("🔗 Тестируем Merge Two Sorted Lists (#21)...")
	fmt.Println("-------------------------------------------")

	testCases := []struct {
		l1       []int
		l2       []int
		expected []int
	}{
		{l1: []int{1, 2, 4}, l2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{l1: []int{}, l2: []int{}, expected: []int{}},
		{l1: []int{}, l2: []int{0}, expected: []int{0}},
		{l1: []int{1, 5, 8}, l2: []int{2, 6}, expected: []int{1, 2, 5, 6, 8}},
	}

	for i, tc := range testCases {
		resList := leetcode.MergeTwoLists(sliceToLinkedList(tc.l1), sliceToLinkedList(tc.l2))
		result := linkedListToSlice(resList)

		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("✅ Тест %d: %v + %v -> %v\n", i+1, tc.l1, tc.l2, result)
		} else {
			fmt.Printf("❌ Тест %d: ОШИБКА! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
	fmt.Println("-------------------------------------------")
}
