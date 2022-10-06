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
func minCameraCover(root *TreeNode) int {
	res := 0
	traverse55(root, false, &res)
	return res
}

// -1: sub node is nil, don't need to be covered
// 0: sub node is not covered
// 1: sub node is covered
// 2: sub node is setted a camera
func traverse55(root *TreeNode, hasParent bool, res *int) int {
	if root == nil {
		return -1
	}
	// 获取左右子节点的情况
	l := traverse55(root.Left, true, res)
	r := traverse55(root.Right, true, res)
	// 根据左右子节点的情况和父节点的情况判断当前节点应该做的事情
	if l == -1 && r == -1 {
		// 当前节点是叶子节点
		if hasParent {
			// 有父节点的话，让父节点来 cover 自己
			return 0
		}
		// 没有父节点的话，自己 set 一个摄像头
		*res++
		return 2
	}
	if l == 0 || r == 0 {
		// 左右子树存在没有被 cover 的
		// 必须在当前节点 set 一个摄像头
		*res++
		return 2
	}
	if l == 2 || r == 2 {
		// 左右子树只要有一个 set 了摄像头
		// 当前节点就已经是 cover 状态了
		return 1
	}
	// 剩下 left == 1 && right == 1 的情况
	// 即当前节点的左右子节点都被 cover
	if hasParent {
		// 如果有父节点的话，可以等父节点 cover 自己
		return 0
	} else {
		// 没有父节点，只能自己 set 一个摄像头
		*res++
		return 2
	}
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func reverse2(b []byte) []byte {
	for i := 0; i < len(b)>>1; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return b
}
