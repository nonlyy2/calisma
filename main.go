package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	fmt.Println("🌳 Тестируем Lowest Common Ancestor (#236)...")
	fmt.Println("-------------------------------------------")

	// Строим дерево:
	//       3
	//      / \
	//     5   1
	//    / \ / \
	//   6  2 0  8
	//     / \
	//    7   4

	root := &leetcode.TreeNode{Val: 3}
	root.Left = &leetcode.TreeNode{Val: 5}
	root.Right = &leetcode.TreeNode{Val: 1}

	node6 := &leetcode.TreeNode{Val: 6}
	node2 := &leetcode.TreeNode{Val: 2}
	root.Left.Left = node6
	root.Left.Right = node2

	node7 := &leetcode.TreeNode{Val: 7}
	node4 := &leetcode.TreeNode{Val: 4}
	node2.Left = node7
	node2.Right = node4

	node5 := root.Left
	node1 := root.Right

	// Тест 1: LCA для 5 и 1 (должен быть 3)
	res1 := leetcode.LowestCommonAncestor2(root, node5, node1)
	check("LCA(5, 1)", res1.Val, 3)

	// Тест 2: LCA для 5 и 4 (должен быть 5, так как 5 - предок 4)
	res2 := leetcode.LowestCommonAncestor2(root, node5, node4)
	check("LCA(5, 4)", res2.Val, 5)

	fmt.Println("-------------------------------------------")
}

func check(name string, got, want int) {
	if got == want {
		fmt.Printf("✅ %s: Получено %d\n", name, got)
	} else {
		fmt.Printf("❌ %s: ОШИБКА! Ожидали %d, получили %d\n", name, want, got)
	}
}
