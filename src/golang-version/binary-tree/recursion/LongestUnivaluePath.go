package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	maxLen(root, root.Val, &res)
	return res
}

// 实现函数的定义：
// 以 root 为根的二叉树从 root 开始值为 parentVal 的最长树枝长度
// 等于左右子树的最长树枝长度的最大值加上 root 节点本身
func maxLen(root *TreeNode, parentVal int, res *int) int {
	if root == nil {
		return 0
	}
	l := maxLen(root.Left, root.Val, res)
	r := maxLen(root.Right, root.Val, res)
	// 后序遍历位置顺便更新全局变量
	// 同值路径就是左右同值树枝长度之和
	*res = max(*res, l+r)
	if root.Val != parentVal {
		return 0
	}
	return max(l, r) + 1
}
