package algorithm

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helperSort(&nums, 0, len(nums))
}

func helperSort(nums *[]int, start, end int) *TreeNode {
	ranges := end - start
	if ranges == 0 {
		return nil
	}
	mid := start + ranges/2
	n := *nums
	root := &TreeNode{Val: n[mid]}
	root.Left = helperSort(nums, start, mid)
	root.Right = helperSort(nums, mid+1, end)
	return root
}
