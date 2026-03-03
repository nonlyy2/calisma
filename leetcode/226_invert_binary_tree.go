package leetcode

// Задача №226: Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}
