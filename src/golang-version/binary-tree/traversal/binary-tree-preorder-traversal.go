package algorithm

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// general template
func preorderTraversal3(node *TreeNode) []int {
	res := []int{}
	preorder(node, &res)
	return res
}

func preorder(node *TreeNode, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
		preorder(node.Left, res)
		preorder(node.Right, res)
	}
}

func preorderTraversal(node *TreeNode) []int {
	res := []int{}
	if node != nil {
		res = append(res, node.Val)
		l := preorderTraversal(node.Left)
		res = append(res, l...)
		r := preorderTraversal(node.Right)
		res = append(res, r...)
	}
	return res
}

// do it iteratively
func preorderTraversal2(node *TreeNode) []int {
	res := []int{}
	if node == nil {
		return res
	}
	stack := []*TreeNode{}
	stack = append(stack, node)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
