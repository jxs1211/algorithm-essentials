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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	l := trimBST(root.Left, low, high)
	r := trimBST(root.Right, low, high)
	if root.Val < low {
		return r // 直接返回 root.right, 等于删除 root 以及 root 的左子树
	}
	if root.Val > high {
		return l // 直接返回 root.left, 等于删除 root 以及 root 的右子树
	}
	root.Left = l
	root.Right = r
	return root
}
