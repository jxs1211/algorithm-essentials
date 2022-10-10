package algorithm

import (
	"strconv"
	"strings"
)

// 297. Serialize and Deserialize Binary Tree
// 449. Serialize and Deserialize BST
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// preorder -> serialize
type Codec struct {
}

func Constructor2() Codec {
	return Codec{}
}

// solution1
// preorder
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, ",")
	return buildPreorder(&s)
	// return buildPostorder(&s)
}

func buildPreorder(datas *[]string) *TreeNode {
	if len(*datas) == 0 {
		return nil
	}
	rootVal := (*datas)[0]
	*datas = (*datas)[1:]
	if rootVal == "#" {
		return nil
	}
	val, _ := strconv.Atoi(rootVal)
	return &TreeNode{
		Val:   val,
		Left:  buildPreorder(datas),
		Right: buildPreorder(datas),
	}
}

// solution2
// postorder
func (this *Codec) serialize2(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return this.serialize2(root.Left) + "," + this.serialize2(root.Right) + "," + strconv.Itoa(root.Val)
}

func (this *Codec) deserialize2(data string) *TreeNode {
	s := strings.Split(data, ",")
	return buildPostorder(&s)
}

//
func buildPostorder(datas *[]string) *TreeNode {
	if len(*datas) == 0 {
		return nil
	}
	ele := (*datas)[len(*datas)-1]
	*datas = (*datas)[:len(*datas)-1]
	if ele == "#" {
		return nil
	}
	val, _ := strconv.Atoi(ele)
	return &TreeNode{
		Val:   val,
		Right: buildPostorder(datas),
		Left:  buildPostorder(datas),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
