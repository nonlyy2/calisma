package main

import (
	"fmt"
	"neetcode/leetcode" // Твой путь к пакету
	"reflect"
)

type TreeNode = leetcode.TreeNode

func main() {
	// ТЕСТ 1: Классическое дерево
	//      3
	//     / \
	//    9  20
	//      /  \
	//     15   7
	// Ожидаем: [[3], [9, 20], [15, 7]]
	root1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	// ТЕСТ 2: Один узел
	// Ожидаем: [[1]]
	root2 := &TreeNode{Val: 1}

	// ТЕСТ 3: Пустое дерево
	// Ожидаем: []
	var root3 *TreeNode = nil

	tests := []struct {
		name     string
		root     *TreeNode
		expected [][]int
	}{
		{"Standard Tree", root1, [][]int{{3}, {9, 20}, {15, 7}}},
		{"Single Node", root2, [][]int{{1}}},
		{"Empty Tree", root3, [][]int{}},
	}

	fmt.Println("🚀 Тестируем Level Order Traversal...")
	for _, tt := range tests {
		result := leetcode.LevelOrder(tt.root)
		if reflect.DeepEqual(result, tt.expected) {
			fmt.Printf("✅ %s: %v\n", tt.name, result)
		} else {
			fmt.Printf("❌ %s: Ожидали %v, получили %v\n", tt.name, tt.expected, result)
		}
	}
}
