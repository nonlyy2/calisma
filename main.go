package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь к пакету
)

type TreeNode = leetcode.TreeNode

func main() {
	// ТЕСТ 1: subRoot действительно является поддеревом
	// root: [3, 4, 5, 1, 2]
	// subRoot: [4, 1, 2]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	subRoot1 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	// ТЕСТ 2: Значения совпадают, но структура поддерева в root шире
	// root: [3, 4, 5, 1, 2, nil, nil, 0] (у узла 2 есть ребенок 0)
	// subRoot: [4, 1, 2]
	// Ожидаем false, так как в основном дереве у "двойки" есть продолжение
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
			},
		},
		Right: &TreeNode{Val: 5},
	}
	subRoot2 := subRoot1 // Те же [4, 1, 2]

	fmt.Println("🌳 Проверка Subtree of Another Tree (#572):")
	fmt.Println("---")
	fmt.Printf("Тест 1 (Является): %v (Ожидаем true)\n", leetcode.IsSubtree(root1, subRoot1))
	fmt.Printf("Тест 2 (Не является): %v (Ожидаем false)\n", leetcode.IsSubtree(root2, subRoot2))
}
