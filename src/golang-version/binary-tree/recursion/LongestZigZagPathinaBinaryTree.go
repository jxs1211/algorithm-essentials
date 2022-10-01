package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	res := 0
	getPathLen(root, &res)
	return res
}

// 输入二叉树的根节点 root，返回两个值
// 第一个是从 root 开始向左走的最长交错路径长度，
// 第一个是从 root 开始向右走的最长交错路径长度
func getPathLen(root *TreeNode, res *int) []int {
	if root == nil {
		return []int{-1, -1}
	}
	l := getPathLen(root.Left, res)
	r := getPathLen(root.Right, res)
	// 后序位置，根据左右子树的交错路径长度推算根节点的交错路径长度
	currLeft := l[1] + 1
	currRight := r[0] + 1
	// 更新全局最大值
	*res = max(*res, max(currLeft, currRight))
	return []int{currLeft, currRight}
}
