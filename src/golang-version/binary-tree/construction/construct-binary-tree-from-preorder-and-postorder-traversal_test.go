package algorithm

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePost(t *testing.T) {
	// Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
	// Output: [1,2,3,4,5,6,7]

	tests := []struct {
		name      string
		preorder  []int
		postorder []int
		want      *TreeNode
	}{
		{
			name:      "base",
			preorder:  []int{1, 2, 4, 5, 3, 6, 7},
			postorder: []int{4, 5, 2, 6, 7, 3, 1},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructFromPrePost(tt.preorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
