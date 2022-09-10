package algorithm

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder(root, k, &res, &count)
	return res
}

func inorder(node *TreeNode, k int, res *int, count *int) {
	if node != nil {
		inorder(node.Left, k, res, count)
		*count = *count + 1
		if k == *count {
			*res = node.Val
			return
		}
		inorder(node.Right, k, res, count)
	}
}
