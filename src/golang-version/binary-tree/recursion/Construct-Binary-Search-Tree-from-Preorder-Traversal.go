package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	i := 1
	for ; i < len(preorder); i++ {
		if preorder[i] > root.Val {
			break
		}
	}
	root.Left = bstFromPreorder(preorder[1:i])
	root.Right = bstFromPreorder(preorder[i:])
	return root
}
