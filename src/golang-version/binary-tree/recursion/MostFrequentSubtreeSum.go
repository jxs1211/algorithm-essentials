package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	res, m := []int{}, make(map[int]int)
	sum(root, m)
	maxCount := 0
	for _, v := range m {
		maxCount = max(v, maxCount)
	}
	for k, _ := range m {
		if m[k] == maxCount {
			res = append(res, k)
		}
	}
	return res
}

func sum(root *TreeNode, m map[int]int) int {
	if root == nil {
		return 0
	}

	l := sum(root.Left, m)
	r := sum(root.Right, m)
	currSum := l + r + root.Val
	m[currSum] = m[currSum] + 1
	return currSum
}
