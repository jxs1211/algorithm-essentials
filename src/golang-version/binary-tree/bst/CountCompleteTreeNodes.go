package algorithm

import "math"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := 0
	left := root
	for left != nil {
		left = left.Left
		l++
	}
	r := 0
	right := root
	for right != nil {
		right = right.Right
		r++
	}
	if l == r {
		return int(math.Pow(2, float64(l)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// count nodes of normal binary tree
func countNormalBT(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNormalBT(root.Left) + countNormalBT(root.Right)
}

// count nodes of full binary tree
func countFullBT(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := 0
	for root != nil {
		root = root.Left
		h++
	}
	return int(math.Pow(2, float64(h)) - 1)
}
