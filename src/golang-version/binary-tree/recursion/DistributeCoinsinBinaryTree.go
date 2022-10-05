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
func distributeCoins(root *TreeNode) int {
	res := 0
	getRest(root, &res)
	return res
}

func getRest(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := getRest(root.Left, res)
	r := getRest(root.Right, res)
	*res += abs2(l) + abs2(r) + (root.Val - 1)
	return l + r + (root.Val - 1)
}

func abs2(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
