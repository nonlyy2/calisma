package leetcode

// Задача №235: Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			break
		}
	}

	return root
}
