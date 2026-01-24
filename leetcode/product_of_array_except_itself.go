package leetcode

// Задача №238: Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	res[0] = 1

	for i := 1; i < length; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	postfix := 1
	for i := length - 1; i >= 0; i-- {
		res[i] = res[i] * postfix
		postfix = postfix * nums[i]
	}

	return res
}
