package algorithm

// reference the illustration of labuladong's answer
// postorder
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	l := root.Left
	r := root.Right

	root.Left = nil
	root.Right = l

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = r
}

// iteration
// Solution2
func flatten2(root *TreeNode) {
	res := preorder(root)
	for i := 1; i < len(res); i++ {
		root.Left = nil
		root.Right = &TreeNode{Val: res[i]}
		root = root.Right
	}
}

func preorder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorder(root.Left)...)
	res = append(res, preorder(root.Right)...)
	return res
}
