package algorithm

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(&nums, 0, len(nums))
}

func helper(nums *[]int, start, end int) *TreeNode {
	ranges := end - start
	if ranges == 0 {
		return nil
	}
	mid := start + ranges/2
	n := *nums
	root := &TreeNode{Val: n[mid]}
	root.Left = helper(nums, start, mid)
	root.Right = helper(nums, mid+1, end)
	return root
}
