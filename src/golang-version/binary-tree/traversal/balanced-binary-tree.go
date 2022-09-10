package algorithm

// Solution1
<<<<<<< HEAD
// postorder
=======
>>>>>>> add and update
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return abs(leftHeight, rightHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}

// Solution2
<<<<<<< HEAD
// postorder
=======
>>>>>>> add and update
func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	is := true
	helper2(root, &is)
	return is
}

func helper2(root *TreeNode, isBlanced *bool) int {
	if root == nil {
		return 0
	}

	l := helper2(root.Left, isBlanced)
	r := helper2(root.Right, isBlanced)
	if abs(l, r) > 1 {
		*isBlanced = false
	}
	return max(l, r) + 1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
