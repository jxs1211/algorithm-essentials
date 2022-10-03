package algorithm

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// postorder
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil {
		return strconv.Itoa(root.Val)
	}
	l := tree2str(root.Left)
	r := tree2str(root.Right)

	if root.Left != nil && root.Right == nil {
		return strconv.Itoa(root.Val) + "(" + l + ")"
	}

	if root.Left == nil && root.Right != nil {
		return strconv.Itoa(root.Val) + "()" + "(" + r + ")"
	}
	return strconv.Itoa(root.Val) + "(" + l + ")" + "(" + r + ")"
}
