package dp

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Solution1
// preorder
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum-root.Val == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// Solution2
// preorder and postorder
func hasPathSum2(root *TreeNode, targetSum int) bool {
	sum := 0
	found := false
	if root == nil {
		return found
	}
	dfs(root, &sum, &found, targetSum)
	return found
}

func dfs(root *TreeNode, sum *int, found *bool, target int) {
	if root == nil {
		return
	}
	*sum += root.Val
	if root.Left == nil && root.Right == nil {
		if *sum == target {
			*found = true
			return
		}
	}
	dfs(root.Left, sum, found, target)
	dfs(root.Right, sum, found, target)
	*sum -= root.Val // avaid dirty data
}
