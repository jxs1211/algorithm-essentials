package algorithm

func zigzagLevelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	nextLayerNodes := []*TreeNode{root}
	traverse1(nextLayerNodes, 0, &res)
	return res
}

func traverse1(layerNodes []*TreeNode, level int, res *[][]int) {
	if len(layerNodes) == 0 {
		return
	}

	currValues := []int{}
	nextLayerNodes := []*TreeNode{}
	for _, node := range layerNodes {
		if level%2 == 0 {
			currValues = append(currValues, node.Val)
		} else {
			currValues = append([]int{node.Val}, currValues...)
		}

		if node.Left != nil {
			nextLayerNodes = append(nextLayerNodes, node.Left)
		}
		if node.Right != nil {
			nextLayerNodes = append(nextLayerNodes, node.Right)
		}
	}
	*res = append(*res, currValues)
	traverse1(nextLayerNodes, level+1, res)
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	traverse3(root, 0, &res)
	return res
}

func traverse3(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) <= level {
		*res = append(*res, []int{})
	}
	if level%2 == 0 {
		(*res)[level] = append((*res)[level], root.Val)
	} else {
		(*res)[level] = append([]int{root.Val}, (*res)[level]...)
	}

	traverse3(root.Left, level+1, res)
	traverse3(root.Right, level+1, res)
}
