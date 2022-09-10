package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 一般来说在找最短路径的时候使用 BFS，其他时候还是 DFS 使用得多一些（主要是递归代码好写）。
// The space complexity of bfs is larger than dfs due to store every layer's nodes in slice
// bfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	q := []*TreeNode{root}
	p := &q
	for len(q) != 0 {
		depth++
		for _ = range *p {
			node := (*p)[0]
			*p = (*p)[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				*p = append(*p, node.Left)
			}
			if node.Right != nil {
				*p = append(*p, node.Right)
			}
		}
	}
	return depth
}

// dfs
func minDepth2(root *TreeNode) int {
	return doMinDepth(root)
}

func doMinDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil {
		return doMinDepth(node.Right) + 1
	}
	if node.Right == nil {
		return doMinDepth(node.Left) + 1
	}
	return min(doMinDepth(node.Left), doMinDepth(node.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
