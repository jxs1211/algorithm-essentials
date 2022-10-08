package algorithm

// 1008. Construct Binary Search Tree from Inorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// postorder
func bstFromInorder(inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	index := len(inorder) >> 1
	rootVal := inorder[index]
	return &TreeNode{
		Val:   rootVal,
		Left:  bstFromInorder(inorder[:index]),
		Right: bstFromInorder(inorder[index+1:]),
	}
}
