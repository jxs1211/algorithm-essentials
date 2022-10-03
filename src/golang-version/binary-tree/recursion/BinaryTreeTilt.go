package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// postorder
func findTilt(root *TreeNode) int {
	res := 0
	dfs111(root, &res)
	return res
}

func dfs111(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	l := dfs111(root.Left, res)
	r := dfs111(root.Right, res)
	*res += abs(l, r)
	return l + r + root.Val
}
