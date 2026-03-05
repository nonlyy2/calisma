package main

import (
	"fmt"
	"neetcode/leetcode"
)

type TreeNode = leetcode.TreeNode

func main() {
	// ТЕСТ 1: Валидное дерево
	//      2
	//     / \
	//    1   3
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}

	// ТЕСТ 2: Невалидное дерево (ловушка!)
	//      5
	//     / \
	//    1   6
	//       / \
	//      4   7
	// Узел 4 находится в правой ветке от 5, значит он ДОЛЖЕН быть больше 5.
	// Но 4 < 5, значит это не BST.
	root2 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 7},
		},
	}

	// ТЕСТ 3: Дерево с дубликатом
	// В правильном BST значения должны быть строго больше/меньше.
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}

	fmt.Println("🌳 Тестируем Validate BST (#98)...")
	fmt.Println("---")
	fmt.Printf("Тест 1 (Валидное):    %v (Ожидаем true)\n", leetcode.IsValidBST(root1))
	fmt.Printf("Тест 2 (Нарушение):   %v (Ожидаем false)\n", leetcode.IsValidBST(root2))
	fmt.Printf("Тест 3 (Дубликат):    %v (Ожидаем false)\n", leetcode.IsValidBST(root3))
}
