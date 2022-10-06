package algorithm

import "reflect"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder
func isCousins(root *TreeNode, x int, y int) bool {
	nodeXParent, nodeYParent := &TreeNode{}, &TreeNode{}
	nodeXDepth, nodeYDepth := 0, 0
	traverse121(root, nil, nodeXParent, nodeYParent, 0, &nodeXDepth, &nodeYDepth, x, y)
	if !reflect.DeepEqual(nodeXParent, nodeYParent) && nodeXDepth == nodeYDepth {
		// 判断 x，y 是否是表兄弟节点
		return true
	}
	return false
}

func traverse121(root, parent, nodeXParent, nodeYParent *TreeNode, level int, nodeXDepth, nodeYDepth *int, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x {
		if parent != nil {
			// 找到 x，记录它的深度和父节点
			nodeXParent.Val, nodeXParent.Left, nodeXParent.Right = parent.Val, parent.Left, parent.Right
			*nodeXDepth = level
		}
	}

	if root.Val == y {
		if parent != nil {
			// 找到 y，记录它的深度和父节点
			nodeYParent.Val, nodeYParent.Left, nodeYParent.Right = parent.Val, parent.Left, parent.Right
			*nodeYDepth = level
		}
	}

	traverse121(root.Left, root, nodeXParent, nodeYParent, level+1, nodeXDepth, nodeYDepth, x, y)
	traverse121(root.Right, root, nodeXParent, nodeYParent, level+1, nodeXDepth, nodeYDepth, x, y)
}
