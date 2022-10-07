package algorithm

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
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
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversal2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inorderTraversal3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal3(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal3() = %v, want %v", got, tt.want)
			}
		})
	}
}
