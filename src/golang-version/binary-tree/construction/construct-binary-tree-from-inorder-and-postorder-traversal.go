package algorithm

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}
	root.Left = buildTree2(inorder[:i], postorder[:i])
	root.Right = buildTree2(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}
