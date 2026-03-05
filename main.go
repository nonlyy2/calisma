package main

import (
	"fmt"
	"neetcode/leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	// Построим BST:
	//        6
	//      /   \
	//     2     8
	//    / \   / \
	//   0   4 7   9
	//      / \
	//     3   5

	n3 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 4, Left: n3, Right: n5}
	n0 := &TreeNode{Val: 0}
	n2 := &TreeNode{Val: 2, Left: n0, Right: n4}

	n7 := &TreeNode{Val: 7}
	n9 := &TreeNode{Val: 9}
	n8 := &TreeNode{Val: 8, Left: n7, Right: n9}

	root := &TreeNode{Val: 6, Left: n2, Right: n8}

	fmt.Println("🌳 Тестируем Lowest Common Ancestor в BST...")

	// Тест 1: LCA для 2 и 8 — это корень (6)
	res1 := leetcode.LowestCommonAncestor(root, n2, n8)
	fmt.Printf("LCA(2, 8): %d (Ожидаем 6)\n", res1.Val)

	// Тест 2: LCA для 2 и 4 — это сам узел 2
	res2 := leetcode.LowestCommonAncestor(root, n2, n4)
	fmt.Printf("LCA(2, 4): %d (Ожидаем 2)\n", res2.Val)
}
