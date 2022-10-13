package algorithm

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	var res, aim *TreeNode
	traverseGetTargetCopy(original, cloned, aim, res)
	return res
}

func traverseGetTargetCopy(original, cloned, target, res *TreeNode) {
	if original == nil || res != nil {
		return
	}
	if original == target {
		res = cloned
		return
	}
	traverseGetTargetCopy(original.Left, cloned.Left, target, res)
	traverseGetTargetCopy(original.Right, cloned.Right, target, res)
}
