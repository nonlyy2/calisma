package main

import (
	"fmt"
	"neetcode/leetcode" // Импортируем твою папку
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// Вызываем функцию через имя пакета
	result := leetcode.TwoSum(nums, target)

	fmt.Println("LeetCode Two Sum:", result)
}
