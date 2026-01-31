package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	// ТЕСТ 1: Покупка за 1, продажа за 6 -> 5
	prices1 := []int{7, 1, 5, 3, 6, 4}
	printTest(prices1, 5, 1)

	// ТЕСТ 2: Цены падают (7->6->4...), прибыли нет -> 0
	prices2 := []int{7, 6, 4, 3, 1}
	printTest(prices2, 0, 2)
}

func printTest(input []int, expected, testNum int) {
	fmt.Printf("--- Test %d ---\n", testNum)
	result := leetcode.MaxProfit(input)
	fmt.Printf("Input: %v | Result: %d (Expected: %d)\n", input, result, expected)
	if result == expected {
		fmt.Println("✅ PASSED")
	} else {
		fmt.Println("❌ FAILED")
	}
	fmt.Println()
}
