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
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	path := []int{}
	traverse4(root, path, &res)
	return res
}

func traverse4(root *TreeNode, path []int, res *[]string) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		s := ""
		for i := 0; i < len(path); i++ {
			s += strconv.Itoa(path[i])
			if i != len(path)-1 {
				s += "->"
			}
		}
		*res = append(*res, s)
		return
	}
	traverse4(root.Left, path, res)
	traverse4(root.Right, path, res)
	path = path[:len(path)-1]
}
