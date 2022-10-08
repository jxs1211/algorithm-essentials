package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Construct Binary Tree from Preorder and Postorder Traversal
// key: which is the index that is the root node of left sub tree in the postorder slice
// postorder
// preorder [root, left, right]
// postorder [left, right, root]
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	preorder, postorder = preorder[1:], postorder[:len(postorder)-1]
	if len(preorder) == 0 {
		return &TreeNode{Val: rootVal}
	}
	i := 0 // the index of the root node of the left sub tree in the postorder slice
	for ; i < len(postorder); i++ {
		if postorder[i] == preorder[0] {
			break
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  constructFromPrePost(preorder[:i+1], postorder[:i+1]),
		Right: constructFromPrePost(preorder[i+1:], postorder[i+1:]),
	}
}
