package algorithm

<<<<<<< HEAD
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
	if l == nil && r == nil {
		return nil
	}
	if l != nil {
		return l
	}
	return r
=======
func lowestCommonAncestor(root *TreeNode, p TreeNode, q TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if max(p.Val, q.Val) < root.Val {
		lowestCommonAncestor(root.Left, p, q)
	}
	if min(p.Val, q.Val) > root.Val {
		lowestCommonAncestor(root.Right, p, q)
	}
	return root
>>>>>>> add and update
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

<<<<<<< HEAD
=======
func preorderTraverse(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraverse(root.Left)...)
	res = append(res, preorderTraverse(root.Right)...)
	return res
}

func inorderTraverse(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraverse(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraverse(root.Right)...)
	return res
}

func postorderTraverse(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, preorderTraverse(root.Left)...)
	res = append(res, preorderTraverse(root.Right)...)
	res = append(res, root.Val)
	return res
}

>>>>>>> add and update
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
