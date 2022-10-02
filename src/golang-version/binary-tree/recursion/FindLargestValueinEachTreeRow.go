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
func largestValues(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	levelEle := []*TreeNode{root}
	for len(levelEle) != 0 {
		p := levelEle[:]
		maxVal := math.MinInt
		for _, node := range p {
			maxVal = max(maxVal, node.Val)
			// update original slice
			levelEle = levelEle[1:]
			if node.Left != nil {
				levelEle = append(levelEle, node.Left)
			}
			if node.Right != nil {
				levelEle = append(levelEle, node.Right)
			}
		}
		res = append(res, maxVal)
	}
	return res
}
