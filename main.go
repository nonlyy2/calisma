package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// Собираем дерево вручную, чтобы проставить Parent
	//        3
	//       / \
	//      5   1
	//     / \
	//    6   2

	node3 := &leetcode.Node{Val: 3}
	node5 := &leetcode.Node{Val: 5, Parent: node3}
	node1 := &leetcode.Node{Val: 1, Parent: node3}
	node6 := &leetcode.Node{Val: 6, Parent: node5}
	node2 := &leetcode.Node{Val: 2, Parent: node5}

	node3.Left = node5
	node3.Right = node1
	node5.Left = node6
	node5.Right = node2

	testCases := []struct {
		p        *leetcode.Node
		q        *leetcode.Node
		expected int
	}{
		{p: node5, q: node1, expected: 3}, // Общий предок - корень
		{p: node5, q: node2, expected: 5}, // p является предком для q
		{p: node6, q: node2, expected: 5}, // Общий предок - 5
	}

	fmt.Println("🌳 Тестируем Lowest Common Ancestor III (#1650)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.LowestCommonAncestor3(tc.p, tc.q)
		if result != nil && result.Val == tc.expected {
			fmt.Printf("Тест %d: ✅ Пройден (LCA: %d)\n", i+1, result.Val)
		} else {
			val := "nil"
			if result != nil {
				val = fmt.Sprintf("%d", result.Val)
			}
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %d, получили %s\n", i+1, tc.expected, val)
		}
	}
}
