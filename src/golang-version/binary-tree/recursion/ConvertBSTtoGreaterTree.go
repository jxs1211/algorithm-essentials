package algorithm

func convertBST(root *TreeNode) *TreeNode {
	res := 0
	rightSum(root, &res)
	return root
}

func rightSum(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	rightSum(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	rightSum(root.Left, sum)
}
