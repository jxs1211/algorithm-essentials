package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	dfs3(root, &res)
	return res
}

func dfs3(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}
	dfs3(root.Left, sum)
	dfs3(root.Right, sum)
}
