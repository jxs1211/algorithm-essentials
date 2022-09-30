package algorithm

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	res := 0
	path := []int{}
	dfsSumNumbers(root, &path, &res)
	return res
}

func dfsSumNumbers(root *TreeNode, path *[]int, res *int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		s := ""
		for _, v := range *path {
			s += strconv.Itoa(v)
		}
		curr, _ := strconv.Atoi(s)
		*res += curr

	}
	dfsSumNumbers(root.Left, path, res)
	dfsSumNumbers(root.Right, path, res)
	*path = (*path)[:len(*path)-1]
}
