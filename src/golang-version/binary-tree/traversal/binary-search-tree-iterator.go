package algorithm

import "container/heap"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postOrder(node *TreeNode, res *IntHeap){
	if node == nil{
		return
	}
	postOrder(node.Left, res)
	postOrder(node.Right, res)
	*res = append(*res, node.Val)
}

type IntHeap []int

func (h IntHeap) Len() int{
	return len(h)
}

func (h IntHeap) Less(i, j int) bool{
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(i interface{}) {
	v := i.(int)
	*h = append(*h, v)
}

func (h *IntHeap) Pop() interface{}{
	n := len(*h)
	ele := (*h)[n - 1]
	*h = (*h)[:n - 1]
	return ele
}

type BSTIterator struct {
    h *IntHeap
}

func Constructor(root *TreeNode) BSTIterator {
    h := &IntHeap{}
	postOrder(root, h)
	heap.Init(h)
	return BSTIterator{
		h: h,
	}
}

func (i *BSTIterator) Next() int {
    return heap.Pop(i.h).(int)
}

func (i *BSTIterator) HasNext() bool {
    return i.h.Len() != 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */