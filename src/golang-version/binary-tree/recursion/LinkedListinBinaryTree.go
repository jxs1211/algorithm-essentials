package algorithm

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TODO
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Input
// [1,4,2,6,8]
// [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
//         1
//     4     4
//  null  2     2  null
//      1 null    6       8
//  null null null null 1    3

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		if check(head, root) {
			return true
		}
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func check(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val == root.Val {
		return check(head.Next, root.Left) || check(head.Next, root.Right)
	}
	return false
}
