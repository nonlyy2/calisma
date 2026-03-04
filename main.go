package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь к пакету
)

type TreeNode = leetcode.TreeNode

func main() {
	// ТЕСТ 1: Классическое дерево
	//      3
	//     / \
	//    9  20
	//      /  \
	//     15   7
	// Ожидаемый результат: 3
	root1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	// ТЕСТ 2: Дерево-палка (вырожденное)
	//  1
	//   \
	//    2
	//     \
	//      3
	// Ожидаемый результат: 3
	root2 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 3},
		},
	}

	// ТЕСТ 3: Пустое дерево
	// Ожидаемый результат: 0
	var root3 *TreeNode = nil

	fmt.Println("🌳 Тестируем MaxDepth...")

	fmt.Printf("Тест 1 (Сбалансированное): %d (Ожидаем 3)\n", leetcode.MaxDepth(root1))
	fmt.Printf("Тест 2 (Палка):            %d (Ожидаем 3)\n", leetcode.MaxDepth(root2))
	fmt.Printf("Тест 3 (Пустое):           %d (Ожидаем 0)\n", leetcode.MaxDepth(root3))
}
