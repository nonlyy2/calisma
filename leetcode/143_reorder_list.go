package leetcode

// Задача №143: Reorder List
// https://leetcode.com/problems/reorder-list/

func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 1st for
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2nd for
	var prev *ListNode
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		tempNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = tempNext
	}

	// 3rd for
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}
