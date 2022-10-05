package algorithm

import (
	"math"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder
type Triple struct {
	node        *TreeNode
	row, column int
}

func doSort(a []*Triple) {
	sort.Slice(a, func(i, j int) bool {
		if a[i].row == a[j].row && a[i].column == a[j].column {
			return a[i].node.Val < a[j].node.Val
		}
		if a[i].column == a[j].column {
			return a[i].row < a[j].row
		}
		return a[i].column < a[j].column
	})
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := []*Triple{}
	res := [][]int{}
	traverse44(root, 0, 0, &nodes)
	doSort(nodes)
	preCol := math.MinInt
	for _, n := range nodes {
		if n.column != preCol {
			res = append(res, []int{})
			preCol = n.column
		}
		res[len(res)-1] = append(res[len(res)-1], n.node.Val)
	}
	return res
}

func traverse44(root *TreeNode, row, col int, res *[]*Triple) {
	if root == nil {
		return
	}
	*res = append(*res, &Triple{root, row, col})
	traverse44(root.Left, row+1, col-1, res)
	traverse44(root.Right, row+1, col+1, res)
}
