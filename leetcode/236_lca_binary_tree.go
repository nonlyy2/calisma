package leetcode

// Задача №236: Lowest Common Ancestor of a Binary Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	left := LowestCommonAncestor2(root.Left, p, q)
	right := LowestCommonAncestor2(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}
