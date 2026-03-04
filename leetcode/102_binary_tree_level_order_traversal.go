package leetcode

// Задача №102: Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	res := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		temp := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			temp[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, temp[:levelSize])
	}
	return res
}
