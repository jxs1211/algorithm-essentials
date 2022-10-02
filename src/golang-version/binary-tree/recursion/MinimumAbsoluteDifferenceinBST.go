package algorithm

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	res := math.MaxInt
	doTraverse(root, prev, &res)
	return res
}

func doTraverse(root, prev *TreeNode, res *int) {
	if root == nil {
		return
	}
	doTraverse(root.Left, prev, res)
	if prev != nil {
		*res = min(*res, abs(root.Val, prev.Val))
	}
	prev = root
	doTraverse(root.Right, prev, res)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
