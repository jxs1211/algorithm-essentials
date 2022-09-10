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

func maxPathSum(root *TreeNode) int {
	res := math.MinInt
	doMaxPathSum(root, &res)
	return res
}

func doMaxPathSum(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := doMaxPathSum(root.Left, res)
	r := doMaxPathSum(root.Right, res)
	currMax := max(l, r) + root.Val
	*res = max(*res, max(currMax, l+r+root.Val))
	return currMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
