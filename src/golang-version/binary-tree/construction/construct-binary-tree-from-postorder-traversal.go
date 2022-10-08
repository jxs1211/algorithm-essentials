package algorithm

// 1008. Construct Binary Search Tree from Postorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// postorder
// [left, right, root]
func bstFromPostorder(postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	index := len(postorder) - 1
	rootVal := postorder[index]
	if len(postorder[:index]) == 0 {
		return &TreeNode{Val: rootVal}
	}
	i := 0
	for ; i < len(postorder); i++ {
		if rootVal < postorder[i] {
			break
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  bstFromPostorder(postorder[:i]),
		Right: bstFromPostorder(postorder[i : len(postorder)-1]),
	}
}
