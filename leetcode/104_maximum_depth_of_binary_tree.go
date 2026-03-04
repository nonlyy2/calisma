package leetcode

// Задача №226: Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}
