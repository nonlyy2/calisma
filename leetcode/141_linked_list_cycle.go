package leetcode

// Задача №141: Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}
