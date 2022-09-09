package algorithm

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	tests := []struct {
		name string
		node *TreeNode
		want []int
	}{
		{
			name: "base",
			node: nil,
			want: []int{},
		},
		{
			name: "node",
			node: &TreeNode{
				Val: 1,
			},
			want: []int{1},
		},
		{
			name: "node with sub node",
			node: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_postorder(t *testing.T) {
	type args struct {
		node *TreeNode
		res  *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postorder(tt.args.node, tt.args.res)
		})
	}
}
