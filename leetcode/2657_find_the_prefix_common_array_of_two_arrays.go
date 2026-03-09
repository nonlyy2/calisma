package leetcode

// Задача №2657: Find the Prefix Common Array of Two Arrays
// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/

func FindThePrefixCommonArray(A []int, B []int) []int {
	seen := make([]int, len(A)+1)
	res := []int{}
	count := 0

	for i := range A {
		seen[A[i]]++
		if seen[A[i]] == 2 {
			count++
		}

		seen[B[i]]++
		if seen[B[i]] == 2 {
			count++
		}

		res = append(res, count)
	}

	return res
}
