package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth <= 1 {
		return &TreeNode{Val: val, Left: root}
	}
	currLevel := 0
	traverse22(root, &currLevel, val, depth)
	return root
}

func traverse22(root *TreeNode, level *int, val int, depth int) {
	if root == nil {
		return
	}
	*level++
	if *level == depth-1 {
		newTreeL := &TreeNode{Val: val, Left: root.Left}
		newTreeR := &TreeNode{Val: val, Right: root.Right}
		root.Left = newTreeL
		root.Right = newTreeR
	}

	traverse22(root.Left, level, val, depth)
	traverse22(root.Right, level, val, depth)
	*level--
}
