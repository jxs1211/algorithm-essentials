package algorithm

// Solution1
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

// Solution2
func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal2(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal2(root.Right)...)
	return res
}

func inorderTraversal3(root *TreeNode) []int {
	res := []int{}
	doInorder(root, &res)
	return res
}

func doInorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	doInorder(root.Left, res)
	*res = append(*res, root.Val)
	doInorder(root.Right, res)
}
