package algorithm

func rightSideView(root *TreeNode) []int {
	res := []int{}
	helper(root, 0, &res)
	return res
}

func helper(node *TreeNode, level int, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	helper(node.Right, level+1, res)
}
