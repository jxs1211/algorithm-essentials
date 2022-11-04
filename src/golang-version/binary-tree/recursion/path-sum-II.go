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

func pathSum(root *TreeNode, targetSum int) [][]int {
	allPaths := [][]int{}
	currentPath := []int{}

	helper(root, currentPath, &allPaths, targetSum)

	return allPaths
}

func helper(root *TreeNode, currentPath []int, allPaths *[][]int, targetSum int) {
	if root == nil {
		return
	}

	newCurrentPath := make([]int, len(currentPath))
	copy(newCurrentPath, currentPath)
	newCurrentPath = append(newCurrentPath, root.Val)

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		*allPaths = append(*allPaths, newCurrentPath)
		return
	}
	helper(root.Left, newCurrentPath, allPaths, targetSum-root.Val)
	helper(root.Right, newCurrentPath, allPaths, targetSum-root.Val)
}
