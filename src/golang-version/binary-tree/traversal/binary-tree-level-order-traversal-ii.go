package algorithm

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	bfs(root, 0, &res)
	reverse(&res)
	return res
}

func reverse(res *[][]int) {
	for i, j := 0, len(*res)-1; i < j; i, j = i+1, j-1 {
		(*res)[i], (*res)[j] = (*res)[j], (*res)[i]
	}
}

// func bfs(node *TreeNode, level int, res *[][]int) {
// 	if node == nil {
// 		return
// 	}
// 	if level+1 > len(*res) {
// 		*res = append(*res, []int{})
// 	}
// 	(*res)[level] = append((*res)[level], node.Val)
// 	bfs(node.Left, level+1, res)
// 	bfs(node.Right, level+1, res)
// }
