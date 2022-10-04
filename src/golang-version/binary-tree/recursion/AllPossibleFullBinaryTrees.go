package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
	m := make(map[int][]*TreeNode, n)
	res := make([]*TreeNode, 0, n)
	if n%2 == 0 {
		return res
	}
	return build(n, m)
}

// 定义：输入一个 n，生成节点树为 n 的所有可能的满二叉树
func build(n int, m map[int][]*TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0, n)
	if n == 1 {
		res = append(res, &TreeNode{})
		return res
	}
	if v, ok := m[n]; ok {
		// 避免冗余计算
		return v
	}
	// 递归生成所有符合条件的左右子树
	for i := 1; i < n; i += 2 {
		j := n - i - 1
		// 利用函数定义，生成左右子树
		l := build(i, m)
		r := build(j, m)
		for _, leftN := range l {
			for _, rightN := range r {
				// 生成根节点,组装出一种可能的二叉树形状,加入结果列表
				res = append(res, &TreeNode{Left: leftN, Right: rightN})
			}
		}
	}
	m[n] = res
	return res
}
