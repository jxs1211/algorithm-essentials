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

func postorderTraverse2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, postorderTraverse2(root.Left)...)
	res = append(res, postorderTraverse2(root.Right)...)
	res = append(res, root.Val)
	return res
}

func postorderTraversal3(node *TreeNode) []int {
	res := []int{}
	postorder(node, &res)
	return res
}

func postorder3(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	postorder3(node.Left, res)
	postorder3(node.Right, res)
	*res = append(*res, node.Val)
}
