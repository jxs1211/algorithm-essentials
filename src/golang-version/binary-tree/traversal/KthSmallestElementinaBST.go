package algorithm

// 230. Kth Smallest Element in a BST
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  inorder
//  the key is the inorder slice which is ascend order, here is the solution of go version
func kthSmallest(root *TreeNode, k int) int {
	res := dfsKthSmallest(root)
	for i, v := range res {
		if i == k-1 {
			return v
		}
	}
	return 0
}

func dfsKthSmallest(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, dfsKthSmallest(root.Left)...)
	res = append(res, root.Val)
	res = append(res, dfsKthSmallest(root.Right)...)
	return res
}
