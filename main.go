package main

import (
	"fmt"
	"neetcode/leetcode" // import
)

func main() {
	nums := []int{2, 7, 11, 15}

	result := leetcode.ContainsDuplicate(nums)

	fmt.Println(result)
}
