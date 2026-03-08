package leetcode

// Задача №1650: Lowest Common Ancestor of a Binary Tree III
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func LowestCommonAncestor3(p *Node, q *Node) *Node {
	a, b := p, q

	for a != b {
		if a == nil {
			a = q
		} else {
			a = a.Parent
		}

		if b == nil {
			b = p
		} else {
			b = b.Parent
		}
	}

	return a
}
