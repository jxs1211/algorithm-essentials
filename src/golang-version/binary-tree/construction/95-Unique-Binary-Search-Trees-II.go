package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  postorder
func generate(lo, hi int) []*TreeNode {
	res := []*TreeNode{}
	if lo > hi {
		res = append(res, nil)
		return res
	}
	for i := lo; i <= hi; i++ {
		l := generate(lo, i-1)
		r := generate(i+1, hi)
		for _, leftNode := range l {
			for _, rightNode := range r {
				root := &TreeNode{Val: i, Left: leftNode, Right: rightNode}
				res = append(res, root)
			}
		}
	}
	return res
}
func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}
