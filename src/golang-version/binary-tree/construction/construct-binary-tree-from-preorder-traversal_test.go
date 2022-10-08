package algorithm

import (
	"reflect"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	// [8,5,1,7,10,12]
	tests := []struct {
		name     string
		preorder []int
		want     *TreeNode
	}{
		{
			name:     "base",
			preorder: []int{8, 5, 1, 7, 10, 12},
			want: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
				Right: &TreeNode{
					Val: 10,
					Right: &TreeNode{
						Val: 12,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Log(inorder11(tt.want))
		t.Log(postorder11(tt.want))
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPreorder(tt.preorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func inorder11(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorder11(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorder11(root.Right)...)
	return res
}

func postorder11(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, postorder11(root.Left)...)
	res = append(res, postorder11(root.Right)...)
	res = append(res, root.Val)
	return res
}
