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
	cal(root, &res)
	return res
}

func cal(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	*res = max(*res, l+r)

	cal(root.Left, res)
	cal(root.Right, res)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	cur := max(l, r) + 1

	return cur
}
