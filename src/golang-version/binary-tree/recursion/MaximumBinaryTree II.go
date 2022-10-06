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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		// 如果 val 是整棵树最大的，那么原来的这棵树应该是 val 节点的左子树，
		// 因为 val 节点是接在原始数组 a 的最后一个元素
		return &TreeNode{Val: val, Left: root}
	}
	// 如果 val 不是最大的，那么就应该在右子树上，
	// 因为 val 节点是接在原始数组 a 的最后一个元素
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}
