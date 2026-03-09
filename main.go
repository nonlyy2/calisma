package main

import (
	"fmt"
	"neetcode/leetcode"
	"reflect"
)

func main() {
	// Вспомогательная функция для создания дерева для тестов
	// Структура: [1,2,3,null,5,null,4]
	root := &leetcode.TreeNode{
		Val: 1,
		Left: &leetcode.TreeNode{
			Val:   2,
			Right: &leetcode.TreeNode{Val: 5},
		},
		Right: &leetcode.TreeNode{
			Val:   3,
			Right: &leetcode.TreeNode{Val: 4},
		},
	}

	testCases := []struct {
		root     *leetcode.TreeNode
		expected []int
	}{
		{root: root, expected: []int{1, 3, 4}},
		{root: nil, expected: []int{}},
		{root: &leetcode.TreeNode{Val: 1, Left: &leetcode.TreeNode{Val: 2}}, expected: []int{1, 2}},
	}

	fmt.Println("🌳 Тестируем Right Side View (#199)...")
	fmt.Println("---")

	for i, tc := range testCases {
		result := leetcode.RightSideView(tc.root)
		if reflect.DeepEqual(result, tc.expected) {
			fmt.Printf("Тест %d: ✅ Пройден. Результат: %v\n", i+1, result)
		} else {
			fmt.Printf("Тест %d: ❌ Ошибка! Ожидали %v, получили %v\n", i+1, tc.expected, result)
		}
	}
}
