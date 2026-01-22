package main

import (
	"fmt"
	"neetcode/leetcode"
)

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Printf("Входные данные: %v\n\n", input)
	result := leetcode.GroupAnagrams(input)
	fmt.Println("Результат (порядок групп может отличаться):")
	fmt.Println("-------------------------------------------")
	for i, group := range result {
		fmt.Printf("Группа %d: %v\n", i+1, group)
	}
}
