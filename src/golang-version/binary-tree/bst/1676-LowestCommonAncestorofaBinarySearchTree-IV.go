package algorithm

func find(root *TreeNode, vals set) *TreeNode {
	if root == nil {
		return nil
	}
	if _, ok := vals[root.Val]; ok {
		return root
	}

	l := find(root.Left, vals)
	r := find(root.Right, vals)
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	return r
}

type set map[int]bool

func lowestCommonAncestorIV(root *TreeNode, nodes []*TreeNode) *TreeNode {
	vals := make(set)
	for _, n := range nodes {
		vals[n.Val] = true
	}
	return find(root, vals)
}
