package leetcode

// Задача №572: Subtree of Another Tree
// https://leetcode.com/problems/subtree-of-another-tree/

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if IsSameTree(root, subRoot) {
		return true
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
