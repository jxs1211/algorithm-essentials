package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	preorder, postorder = preorder[1:], postorder[:len(postorder)-1]
	if len(preorder) == 0 {
		return root
	}
	i := 0
	for ; i < len(postorder); i++ {
		if postorder[i] == preorder[0] {
			break
		}
	}

	root.Left = constructFromPrePost(preorder[:i+1], postorder[:i+1])
	root.Right = constructFromPrePost(preorder[i+1:], postorder[i+1:])
	return root
}
