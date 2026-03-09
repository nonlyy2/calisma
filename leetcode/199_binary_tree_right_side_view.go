package leetcode

// Задача №199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/

func RightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{}
	queue = append(queue, root)
	res := []int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		res = append(res, queue[(levelSize-1)].Val)

		for i := range levelSize {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[levelSize:]
	}

	return res
}
