package algorithm

import "math"

func isValidBST(root *TreeNode) bool {

	return helper(root, math.Inf(-1), math.Inf(1))
}

func helper(node *TreeNode, min float64, max float64) bool {
	if node == nil {
		return true
	}
	v := float64(node.Val)
	return helper(node.Left, min, v) && helper(node.Right, v, max) && min < v && v < max
}
