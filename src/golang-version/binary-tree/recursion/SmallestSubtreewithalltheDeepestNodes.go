package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// postorder
type Result struct {
	node  *TreeNode // 记录最近公共祖先节点 node
	depth int       // 记录以 node 为根的二叉树最大深度
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	return traverse11(root).node
}

// 定义：输入一棵二叉树，返回该二叉树的最大深度以及最深叶子节点的最近公共祖先节点
func traverse11(root *TreeNode) *Result {
	if root == nil {
		return &Result{nil, 0}
	}
	leftRes := traverse11(root.Left)
	rightRes := traverse11(root.Right)
	if leftRes.depth == rightRes.depth {
		// 当左右子树的最大深度相同时，这个根节点是新的最近公共祖先
		// 以当前 root 节点为根的子树深度是子树深度 + 1
		return &Result{node: root, depth: leftRes.depth + 1}
	}
	// 如果不相等，最深叶子节点的最近公共祖先一定在更深的子树
	if leftRes.depth > rightRes.depth {
		leftRes.depth++
		return leftRes
	}
	rightRes.depth++
	return rightRes
}
