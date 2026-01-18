package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := leetcode.TwoSumII(numbers, target)

	fmt.Println(result)
}
