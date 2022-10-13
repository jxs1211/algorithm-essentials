package algorithm

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  postorder
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	dfsMaxAncestorDiff(root, &res)
	return res
}

func dfsMaxAncestorDiff(root *TreeNode, res *int) []int {
	if root == nil {
		return []int{math.MaxInt, math.MinInt}
	}
	l := dfsMaxAncestorDiff(root.Left, res)
	r := dfsMaxAncestorDiff(root.Right, res)
	currMin := min(root.Val, min(l[0], r[0]))
	currMax := max(root.Val, max(l[1], r[1]))
	*res = max(*res, max(currMax-root.Val, root.Val-currMin))
	return []int{currMin, currMax}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func abs(a, b int) int {
// 	if a > b {
// 		return a - b
// 	}
// 	return b - a
// }

// func dfs(root *TreeNode, res *int) int {
//     if root == nil {return -1}
//     l := dfs(root.Left, res)
//     r := dfs(root.Right, res)
//     // if l != -1 && r != -1{
//     //     *res = max(*res, max(abs(root.Val, l), abs(root.Val, r)))
//     // }
//     if l != -1{
//         *res = max(*res, abs(root.Val, l))
//     }
//     if r != -1{
//         *res = max(*res, abs(root.Val, r))
//     }
//     return root.Val
// }
