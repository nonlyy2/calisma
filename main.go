package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	s := "anagram"
	t := "nagaram"

	result := leetcode.IsAnagram(s, t)

	fmt.Println(result)
}
