package leetcode

import "strconv"

// Задача №443: String Compression
// https://leetcode.com/problems/string-compression/

func Compress(chars []byte) int {
	write := 0
	i := 0

	for i < len(chars) {
		ch := chars[i]
		count := 0

		for i < len(chars) && chars[i] == ch {
			count++
			i++
		}

		chars[write] = ch
		write++

		if count > 1 {
			countStr := strconv.Itoa(count)
			for j := 0; j < len(countStr); j++ {
				chars[write] = countStr[j]
				write++
			}
		}
	}

	return write
}
