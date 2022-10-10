package algorithm

// 538. Convert BST to Greater Tree
// 1038. Binary Search Tree to Greater Sum Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	traverseConvertBST(root, &sum)
	return root
}

func traverseConvertBST(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traverseConvertBST(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traverseConvertBST(root.Left, sum)
}

// descend slice of the tree node
func inorderDescend(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderDescend(root.Right)...)
	res = append(res, root.Val)
	res = append(res, inorderDescend(root.Left)...)
	return res
}
