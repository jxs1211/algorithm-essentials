package algorithm

func inorderTraversal(node *TreeNode) []int {
	res := []int{}
	inorder(node, &res)
	return res
}

func inorder(node *TreeNode, res *[]int) {
	if node != nil {
		inorder(node.Left, res)
		*res = append(*res, node.Val)
		inorder(node.Right, res)
	}
}
