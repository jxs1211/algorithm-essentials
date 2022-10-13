package algorithm

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	res := 1
	firstId := []int{}
	doWidthOfBinaryTree(root, 1, 1, &res, &firstId)
	return res
}

func doWidthOfBinaryTree(root *TreeNode, id, depth int, res *int, firstId *[]int) {
	if root == nil {
		return
	}

	if len(*firstId) == depth-1 {
		//only left sub node will be added to the slice
		*firstId = append(*firstId, id)
	} else {
		// calculate diff of the current id and the first id(id of left sub node in the same depth)
		*res = max(*res, id-(*firstId)[depth-1]+1)
	}
	doWidthOfBinaryTree(root.Left, id*2, depth+1, res, firstId)
	doWidthOfBinaryTree(root.Right, id*2+1, depth+1, res, firstId)
}
