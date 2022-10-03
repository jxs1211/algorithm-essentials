package algorithm

type Node1 struct {
	Val      int
	Children []*Node1
}

// postorder
// Solution
func maxDepth2(root *Node1) int {
	currMax, res := 0, 0
	dfs22(root, &currMax, &res)
	return res
}

func dfs22(root *Node1, currMax, res *int) {
	if root == nil {
		return
	}
	*currMax++
	*res = max(*res, *currMax)
	for _, node := range root.Children {
		dfs22(node, currMax, res)
	}
	*currMax--
}

// Solution
func maxDepth4(root *Node1) int {
	res := 0
	dfs4(root, &res)
	return res
}

func dfs4(root *Node1, res *int) int {
	if root == nil {
		return 0
	}
	currMax := 0
	for _, node := range root.Children {
		currMax = max(currMax, dfs4(node, res))
	}

	*res = max(*res, currMax+1)
	return currMax + 1
}
