package leetcode

import "math"

// Задача №98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/

func IsValidBST(root *TreeNode) bool {
	return ValidateBST(root, math.MinInt64, math.MaxInt64)
}

func ValidateBST(node *TreeNode, low, high int) bool {
	if node == nil {
		return true
	}

	if node.Val <= low || node.Val >= high {
		return false
	}

	return ValidateBST(node.Left, low, node.Val) && ValidateBST(node.Right, node.Val, high)
}
