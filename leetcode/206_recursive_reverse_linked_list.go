package leetcode

// Задача №206: Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/

func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseListRecursive(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}
