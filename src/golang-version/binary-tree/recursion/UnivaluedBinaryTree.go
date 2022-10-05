package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// preorder
func isUnivalTree(root *TreeNode) bool {
	res, preVal := true, root.Val
	if root == nil {
		return res
	}
	traverse66(root, &preVal, &res)
	return res
}

func traverse66(root *TreeNode, preVal *int, res *bool) {
	if root == nil || !*res {
		return
	}
	if root.Val != *preVal {
		*res = false
		return
	}
	traverse66(root.Left, preVal, res)
	traverse66(root.Right, preVal, res)
}
