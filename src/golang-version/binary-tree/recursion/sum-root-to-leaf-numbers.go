package algorithm

<<<<<<< HEAD
import "strconv"

=======
>>>>>>> add and update
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
<<<<<<< HEAD
	res := 0
	path := []int{}
	dfsSumNumbers(root, &path, &res)
	return res
}

func dfsSumNumbers(root *TreeNode, path *[]int, res *int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root.Left == nil && root.Right == nil {
		s := ""
		for _, v := range *path {
			s += strconv.Itoa(v)
		}
		curr, _ := strconv.Atoi(s)
		*res += curr

	}
	dfsSumNumbers(root.Left, path, res)
	dfsSumNumbers(root.Right, path, res)
	*path = (*path)[:len(*path)-1]
=======
	return dfs(root, 0)
}

func dfs(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}
	curr := sum*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return curr
	}
	return dfs(node.Left, curr) + dfs(node.Right, curr)
>>>>>>> add and update
}
