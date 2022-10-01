package algorithm

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	var res, aim *TreeNode
	traverse(original, cloned, aim, res)
	return res
}

func traverse(original, cloned, target, res *TreeNode) {
	if original == nil || res != nil {
		return
	}
	if original == target {
		res = cloned
		return
	}
	traverse(original.Left, cloned.Left, target, res)
	traverse(original.Right, cloned.Right, target, res)
}
