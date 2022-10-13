package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
	*TreeNode
	set map[int]bool
}

func Constructor3(root *TreeNode) FindElements {
	set := make(map[int]bool)
	traverse34(root, 0, set)
	return FindElements{
		TreeNode: root,
		set:      set,
	}
}

func traverse34(root *TreeNode, val int, set map[int]bool) {
	if root == nil {
		return
	}
	root.Val = val
	set[val] = true
	traverse34(root.Left, 2*val+1, set)
	traverse34(root.Right, 2*val+2, set)
}

func (this *FindElements) Find(target int) bool {
	_, ok := this.set[target]
	return ok
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
