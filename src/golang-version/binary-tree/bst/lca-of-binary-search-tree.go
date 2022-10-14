package algorithm

// postorder
func lowestCommonAncestorBinaryTree2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestorBinaryTree2(root.Left, p, q)
	r := lowestCommonAncestorBinaryTree2(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
