package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	dfsPathSum(root, sum, &res, []int{})
	return res
}

func dfsPathSum(root *TreeNode, sum int, res *[][]int, path []int) {
	if root == nil {
		return
	}
	remain := sum - root.Val
	path = append(path, root.Val)
	if remain == 0 && root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	dfsPathSum(root.Left, remain, res, path)
	dfsPathSum(root.Right, remain, res, path)
}
