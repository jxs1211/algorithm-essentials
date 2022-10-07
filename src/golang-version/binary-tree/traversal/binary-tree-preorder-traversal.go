package algorithm

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// general template
func preorderTraversal3(node *TreeNode) []int {
	res := []int{}
	preorder3(node, &res)
	return res
}

func preorder3(node *TreeNode, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
		preorder3(node.Left, res)
		preorder3(node.Right, res)
	}
}

func preorderTraverse2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, preorderTraverse2(root.Left)...)
	res = append(res, preorderTraverse2(root.Right)...)
	return res
}

// do it iteratively
func preorderTraversal4(node *TreeNode) []int {
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
