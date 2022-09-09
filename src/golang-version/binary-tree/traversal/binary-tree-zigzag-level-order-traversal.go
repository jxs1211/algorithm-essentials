package algorithm

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	helper2(root, 0, &res)
	return res
}

func helper2(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	if len(*res) < level+1 {
		*res = append(*res, []int{})
	}
	if level%2 == 0 {
		(*res)[level] = append((*res)[level], node.Val)
	} else {
		(*res)[level] = append([]int{node.Val}, (*res)[level]...)
	}
	helper2(node.Left, level+1, res)
	helper2(node.Right, level+1, res)
}
