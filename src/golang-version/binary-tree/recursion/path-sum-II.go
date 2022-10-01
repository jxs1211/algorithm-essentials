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
	dfs(root, sum, &res, []int{})
	return res
}

func dfs(root *TreeNode, sum int, res *[][]int, path []int) {
	if root == nil {
		return
	}
	remain := sum - root.Val
	path = append(path, root.Val)
	if remain == 0 && root.Left == nil && root.Right == nil {
		*res = append(*res, path)
		return
	}
	dfs(root.Left, remain, res, path)
	dfs(root.Right, remain, res, path)
}
