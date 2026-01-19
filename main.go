package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("input: %v\n", nums1)
	res1 := leetcode.ThreeSum(nums1)
	fmt.Printf("result: %v\n", res1)
	fmt.Println("expected: [[-1 -1 2] [-1 0 1]]")
}
