package algorithm

import "strings"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	res := ""
	path := []byte{}
	traverse333(root, path, &res)
	return res
}

func traverse333(root *TreeNode, path []byte, res *string) {
	if root == nil {
		return
	}
	path = append(path, 'a'+byte(root.Val))
	if root.Left == nil && root.Right == nil {
		// path = reverse(path)
		reversedPath := reverseAndCopy(path)
		s := string(reversedPath)
		if len(*res) == 0 || strings.Compare(s, *res) < 0 {
			*res = s
		}
		// path = reverse(path)
		path = path[:len(path)-1]
		return
	}

	traverse333(root.Left, path, res)
	traverse333(root.Right, path, res)
	path = path[:len(path)-1]
}

func reverseAndCopy(b []byte) []byte {
	cop := make([]byte, len(b))
	for i := len(b) - 1; i >= 0; i-- {
		j := len(b) - i - 1
		cop[j] = b[i]
	}
	return cop
}
