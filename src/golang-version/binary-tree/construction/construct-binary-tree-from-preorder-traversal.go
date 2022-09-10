package algorithm

// 1008. Construct Binary Search Tree from Preorder Traversal
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
<<<<<<< HEAD
// Given an array of integers preorder, which represents the preorder traversal of a BST (i.e., binary search tree), construct the tree and return its root.

// It is guaranteed that there is always possible to find a binary search tree with the given requirements for the given test cases.

// A binary search tree is a binary tree where for every node, any descendant of Node.left has a value strictly less than Node.val, and any descendant of Node.right has a value strictly greater than Node.val.

// A preorder traversal of a binary tree displays the value of the node first, then traverses Node.left, then traverses Node.right.

// postorder
=======
>>>>>>> add and update
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
