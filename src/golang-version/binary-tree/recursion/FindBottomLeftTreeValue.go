package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	res, currLevel, maxLevel := 0, 0, 0
	doTraverse2(root, &res, &currLevel, &maxLevel)
	return res
}

func doTraverse2(root *TreeNode, res, currLevel, maxLevel *int) {
	if root == nil {
		return
	}
	*currLevel++
	if *currLevel > *maxLevel { //at the first time dive into a new layer
		*maxLevel = *currLevel
		*res = root.Val
	}
	doTraverse2(root.Left, res, currLevel, maxLevel)
	doTraverse2(root.Right, res, currLevel, maxLevel)
	*currLevel--
}
