package algorithm

// Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestor(root *TreeNode, p TreeNode, q TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if max(p.Val, q.Val) < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if min(p.Val, q.Val) > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// find the target
	return root
}
