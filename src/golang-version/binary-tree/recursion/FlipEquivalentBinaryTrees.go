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
// 定义：输入两棵二叉树，判断这两棵二叉树是否是翻转等价的
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	// 判断 root1 和 root2 两个节点是否能够匹配
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	// 判断子树是否能够匹配，不翻转、翻转两种情况满足一种即可算是匹配
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) || (flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}
