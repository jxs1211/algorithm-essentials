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
// Solution1
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := increasingBST(root.Left)
	root.Left = nil
	r := increasingBST(root.Right)
	root.Right = r
	if l == nil {
		return root
	}
	p := l
	for p != nil && p.Right != nil {
		p = p.Right
	}
	p.Right = root
	return l
}

// preorder + postorder with inorder slice
// Solution2
func increasingBST2(root *TreeNode) *TreeNode {
	inorder := []int{}
	traverse3(root, &inorder)
	return buildTree(inorder)
}

func buildTree(inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: inorder[0]}
	root.Right = buildTree(inorder[1:])
	return root
}

func traverse3(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traverse3(root.Left, res)
	*res = append(*res, root.Val)
	traverse3(root.Right, res)
}
