package leetcode

// Задача №2: Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}

		curr = curr.Next
	}

	return dummy.Next
}
