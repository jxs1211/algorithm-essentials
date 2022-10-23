package algorithm

import (
	"reflect"
	"testing"
)

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
			if got := preorderTraverse2(tt.node); !reflect.DeepEqual(got, tt.want) {
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

func Test_preorderTraversal4(t *testing.T) {
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
			if got := preorderTraversal4(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal3() = %v, want %v", got, tt.want)
			}
		})
	}
}
