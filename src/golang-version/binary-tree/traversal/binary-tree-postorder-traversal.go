package algorithm

func postorderTraversal(node *TreeNode) []int {
	res := []int{}
	postorder(node, &res)
	return res
}

func postorder(node *TreeNode, res *[]int) {
	if node != nil {
		postorder(node.Left, res)
		postorder(node.Right, res)
		*res = append(*res, node.Val)
	}
}
