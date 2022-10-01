package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Input:

func maxProduct(root *TreeNode) int {
	res, treeSum := 0, 0
	treeSum = getSum(root, &res, &treeSum)
	getSum(root, &res, &treeSum)
	return int(res % (1e9 + 7))
}

func getSum(root *TreeNode, res *int, treeSum *int) int {
	if root == nil {
		return 0
	}
	l := getSum(root.Left, res, treeSum)
	r := getSum(root.Right, res, treeSum)
	rootSum := root.Val + l + r
	*res = max(*res, rootSum*(*treeSum-rootSum))
	return rootSum
}
