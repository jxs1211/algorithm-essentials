package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bstFromInorder(t *testing.T) {
	tests := []struct {
		name    string
		inorder []int
		want    *TreeNode
	}{
		{
			name:    "case1: 10 is root and 12 is right sub node",
			inorder: []int{1, 5, 7, 8, 10, 12},
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
		{
			name:    "case2: 12 is root and 10 is left sub node",
			inorder: []int{1, 5, 7, 8, 10, 12},
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
					Val: 12,
					Left: &TreeNode{
						Val: 10,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bstFromInorder(tt.inorder)
			assert.Equal(t, tt.want, got, tt.name)
		})
	}
}
