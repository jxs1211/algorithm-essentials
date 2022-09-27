package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	curr := sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return curr
	}
	return dfs(node.Left, curr) + dfs(node.Right, curr)
}
