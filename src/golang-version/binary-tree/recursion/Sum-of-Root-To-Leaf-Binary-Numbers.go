package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	res, path := 0, 0
	dfsSumRootToLeaf(root, &res, &path)
	return res
}

func dfsSumRootToLeaf(root *TreeNode, res *int, path *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*res += *path<<1 | root.Val
	}
	*path = *path<<1 | root.Val
	dfsSumRootToLeaf(root.Left, res, path)
	dfsSumRootToLeaf(root.Right, res, path)
	*path = *path >> 1
}
