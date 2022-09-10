package algorithm

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sl := []string{}
	res := this.doSerialize(root, sl)
	return strings.Join(res, " ")
}

func (this *Codec) doSerialize(node *TreeNode, s []string) []string {
	if node == nil {
		s = append(s, "#")
		return s
	}
	s = append(s, strconv.Itoa(node.Val))
	s = this.doSerialize(node.Left, s)
	s = this.doSerialize(node.Right, s)
	return s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data, " ")
	return this.doDeserialize(&s)
}

func (this *Codec) doDeserialize(data *[]string) *TreeNode {
	if len(*data) == 0 {
		return nil
	}
	ele := (*data)[0]
	if ele == "#" {
		*data = (*data)[1:]
		return nil
	}
	v, _ := strconv.Atoi(ele)
	root := &TreeNode{Val: v}
	*data = (*data)[1:]
	root.Left = this.doDeserialize(data)
	root.Right = this.doDeserialize(data)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
