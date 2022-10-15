package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	diameter(root, &res)
	return res
}

func diameter(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := diameter(root.Left, res)
	r := diameter(root.Right, res)
	*res = max(*res, l+r)
	return 1 + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
