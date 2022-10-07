package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	set := map[int]bool{}
	res := []*TreeNode{}
	for _, i := range to_delete {
		set[i] = true
	}
	delete(root, false, set, &res)
	return res
}

func delete(root *TreeNode, hasParent bool, set map[int]bool, res *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	_, toDelete := set[root.Val]
	if !toDelete && !hasParent {
		*res = append(*res, root)
	}

	root.Left = delete(root.Left, !toDelete, set, res)
	root.Right = delete(root.Right, !toDelete, set, res)

	if toDelete {
		return nil
	}
	return root
}
