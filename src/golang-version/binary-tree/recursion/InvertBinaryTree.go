package algorithm

// solution1
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := invertTree(root.Left)
	r := invertTree(root.Right)
	root.Left = r
	root.Right = l
	return root
}

// solution2
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree2(root.Left)
	invertTree2(root.Right)
	return root
}
