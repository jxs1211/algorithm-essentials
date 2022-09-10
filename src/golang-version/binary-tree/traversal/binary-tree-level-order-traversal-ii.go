package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	nextNodes := []*TreeNode{root}
	traverse2(nextNodes, &res)
	return res
}

func traverse2(nextNodes []*TreeNode, res *[][]int) {
	if len(nextNodes) == 0 {
		return
	}
	currValues := []int{}
	nextLayerNodes := []*TreeNode{}
	for _, node := range nextNodes {
		currValues = append(currValues, node.Val)
		if node.Left != nil {
			nextLayerNodes = append(nextLayerNodes, node.Left)
		}
		if node.Right != nil {
			nextLayerNodes = append(nextLayerNodes, node.Right)
		}
	}

	traverse(nextLayerNodes, res)
	*res = append(*res, currValues)
}
