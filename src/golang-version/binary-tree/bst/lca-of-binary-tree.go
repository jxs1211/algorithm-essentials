package algorithm

// Lowest Common Ancestor of a Binary Tree
// postorder
func lowestCommonAncestorBinaryTree(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	l := lowestCommonAncestorBinaryTree(root.Left, p, q)
	r := lowestCommonAncestorBinaryTree(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}
