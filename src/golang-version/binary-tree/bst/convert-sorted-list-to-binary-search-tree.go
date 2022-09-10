/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package algorithm


func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil{
        return nil
    }
    if head.Next == nil{
        return &TreeNode{Val: head.Val}
    }
    mid := cutMiddle(head)
    root := &TreeNode{Val: mid.Val}
    root.Left = sortedListToBST(head)
    root.Right = sortedListToBST(mid.Next)
    return root
}

func cutMiddle(node *ListNode) *ListNode{
    if node == nil {
        return nil
    }
    slow, fast, pre := node, node, node
    for fast != nil && fast.Next != nil{
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    pre.Next = nil
    return slow
}