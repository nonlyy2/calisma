package leetcode

// Задача №238: Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/

func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	prefix := make([]int, length)
	suffix := make([]int, length)

	for i := 0; i < length-1; i++ {
		if i == 0 {
			prefix[0] = 1
			continue
		}
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for j := length - 1; j >= 0; j-- {
		if j == length-1 {
			suffix[length-1] = 1
			continue
		}
		suffix[j] = suffix[j+1] * nums[j+1]
	}

	for k := 0; k < length; k++ {
		res[k] = prefix[k] * suffix[k]
	}

	return res
}
