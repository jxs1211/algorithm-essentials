package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  preorder
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left != nil && root.Right != nil {
			// 1 record minNode
			minNode := getMinNode(root.Right)
			// 2 delete minNode and update root.Right
			root.Right = deleteNode(root.Right, minNode.Val)
			// 3 replace root with minNode
			minNode.Left, minNode.Right = root.Left, root.Right
			// root.Val = r.Val
			root = minNode
		}
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func getMinNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return getMinNode(root.Left)
}

func getMinNode2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}
