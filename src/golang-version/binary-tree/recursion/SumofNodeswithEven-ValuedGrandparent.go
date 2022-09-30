package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	res := 0
	dfsSumEvenGrandparent(root, &res)
	return res
}

func dfsSumEvenGrandparent(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				*res += root.Left.Left.Val
			}
			if root.Left.Right != nil {
				*res += root.Left.Right.Val
			}
		}
		if root.Right != nil {
			if root.Right.Left != nil {
				*res += root.Right.Left.Val
			}
			if root.Right.Right != nil {
				*res += root.Right.Right.Val
			}
		}
	}
	dfsSumEvenGrandparent(root.Left, res)
	dfsSumEvenGrandparent(root.Right, res)
}
