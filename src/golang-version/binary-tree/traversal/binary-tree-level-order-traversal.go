package algorithm

//Solution1
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	nextNodes := []*TreeNode{root}
	traverse(nextNodes, &res)
	return res
}

func traverse(nextNodes []*TreeNode, res *[][]int) {
	if len(nextNodes) == 0 {
		return
	}
	currValues := []int{}
	nextLayerNodes := []*TreeNode{}
	for _, node := range nextNodes {
		currValues = append(currValues, node.Val)
		if node.Left != nil {
			nextLayerNodes = append(nextLayerNodes, node.Left)
		}
		if node.Right != nil {
			nextLayerNodes = append(nextLayerNodes, node.Right)
		}
	}

	*res = append(*res, currValues)
	traverse(nextLayerNodes, res)
}

//Solution2
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)
}

//Solution3
func levelOrder3(root *TreeNode) [][]int {
	res := [][]int{}
	q := []*TreeNode{root}
	if root == nil {
		return res
	}
	qPtr := &q
	for len(*qPtr) != 0 {
		curr := []int{}
		cp := (*qPtr)[:]
		for i := 0; i < len(cp); i++ {
			currNode := cp[i]
			curr = append(curr, currNode.Val)
			if currNode.Left != nil {
				*qPtr = append(*qPtr, currNode.Left)
			}
			if currNode.Right != nil {
				*qPtr = append(*qPtr, currNode.Right)
			}
			*qPtr = (*qPtr)[1:]
		}
		res = append(res, curr)
	}
	return res
}
