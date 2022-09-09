package algorithm

//
func levelOrder(node *TreeNode) [][]int {
	res := [][]int{}
	bfs(node, 0, &res)
	return res
}

func bfs(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	if level+1 > len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], node.Val)
	bfs(node.Left, level+1, res)
	bfs(node.Right, level+1, res)
}
