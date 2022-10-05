package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	canFlip, res, i := true, []int{}, 0
	// try to flip in the process of traversing
	traverse33(root, &canFlip, &i, &voyage, &res)
	if canFlip {
		return res
	}
	return []int{-1}
}

func traverse33(root *TreeNode, canFlip *bool, i *int, voyage *[]int, res *[]int) {
	if root == nil || !*canFlip {
		return
	}
	// don't need to flip any more if the values of root and the corresponding value in voyage don't equal
	if root.Val != (*voyage)[*i] {
		*canFlip = false
		return
	}
	*i++
	if root.Left != nil && root.Left.Val != (*voyage)[*i] {
		// try to flip the left sub tree and right sub tree if the result of the preorder traversing doesn't match the voyage
		root.Left, root.Right = root.Right, root.Left
		// record the flipped node's value
		*res = append(*res, root.Val)
	}
	traverse33(root.Left, canFlip, i, voyage, res)
	traverse33(root.Right, canFlip, i, voyage, res)
}
