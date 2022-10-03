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
// postorder
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxVal := math.MinInt
	index := 0
	for i := 0; i < len(nums); i++ {
		if maxVal < nums[i] {
			index = i
			maxVal = nums[i]
		}
	}
	return &TreeNode{
		Val:   maxVal,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
}
