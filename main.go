package main

import (
	"fmt"
	"neetcode/leetcode" // Замени на свой путь к пакету
)

type TreeNode = leetcode.TreeNode

func main() {
	// ТЕСТ 1: Идентичные деревья [1, 2, 3]
	p1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	// ТЕСТ 2: Разная структура [1, 2] и [1, nil, 2]
	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	// ТЕСТ 3: Одинаковая структура, разные значения [1, 2, 1] и [1, 1, 2]
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}

	fmt.Println("🌳 Проверка Same Tree (#100):")
	fmt.Println("---")
	fmt.Printf("Тест 1 (Идентичные): %v (Ожидаем true)\n", leetcode.IsSameTree(p1, q1))
	fmt.Printf("Тест 2 (Структура):  %v (Ожидаем false)\n", leetcode.IsSameTree(p2, q2))
	fmt.Printf("Тест 3 (Значения):   %v (Ожидаем false)\n", leetcode.IsSameTree(p3, q3))
}
