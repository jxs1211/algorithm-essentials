package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  preorder\inorder\postorder are all ok
// Solution1
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	sumRangeSumBST(root, low, high, &res)
	return res
}

func sumRangeSumBST(root *TreeNode, low, high int, res *int) {
	if root == nil {
		return
	}
	if root.Val >= low && root.Val <= high {
		*res += root.Val
	}
	sumRangeSumBST(root.Left, low, high, res)
	sumRangeSumBST(root.Right, low, high, res)
}

// Solution2
func rangeSumBST2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val >= low && root.Val <= high {
		return root.Val + rangeSumBST2(root.Left, low, high) + rangeSumBST2(root.Right, low, high)
	} else if root.Val < low {
		return rangeSumBST2(root.Right, low, high)
	} else {
		return rangeSumBST2(root.Left, low, high)
	}
}
