package leetcode

// Задача №271: Encode and Decode Strings
// https://leetcode.com/problems/encode-and-decode-strings/

import (
	"fmt"
	"strconv"
	"strings"
)

func Encode(strs []string) string {
	var builder strings.Builder
	for _, val := range strs {
		fmt.Fprintf(&builder, "%d", len(val))
		builder.WriteByte('#')
		builder.WriteString(val)
	}

	return builder.String()
}

func Decode(s string) []string {
	result := []string{}
	i := 0

	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(s[i:j])
		start := j + 1
		end := start + length

		result = append(result, s[start:end])
		i = end
	}

	return result
}
