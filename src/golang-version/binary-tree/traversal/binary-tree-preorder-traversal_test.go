package algorithm

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
			name: "one node",
			node: &TreeNode{
				Val: 1,
			},
			want: []int{1},
		},
		{
			name: "one node with sub node",
			node: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal2(t *testing.T) {
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
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal2(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorderTraversal3(t *testing.T) {
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
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal3(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preorder(t *testing.T) {
	type args struct {
		node *TreeNode
		res  *[]int
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preorder(tt.args.node, tt.args.res)
		})
	}
}
