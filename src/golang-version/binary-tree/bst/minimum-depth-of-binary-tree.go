package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// update count for current layer and dive into the other layer if the current one is nil

func minDepth(root *TreeNode) int {
	return minDepthHelper(root)
}

func minDepthHelper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil {
		return minDepthHelper(node.Right) + 1
	}
	if node.Right == nil {
		return minDepthHelper(node.Left) + 1
	}
	return min(minDepthHelper(node.Left), minDepthHelper(node.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
