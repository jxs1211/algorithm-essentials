package algorithm

import (
	"reflect"
	"testing"
)

func Test_bstFromPostorder(t *testing.T) {
	tests := []struct {
		name      string
		postorder []int
		want      *TreeNode
	}{
		{
			name:      "nil tree",
			postorder: []int{},
			want:      nil,
		},
		{
			name:      "real tree",
			postorder: []int{1, 7, 5, 12, 10, 8},
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
		t.Run(tt.name, func(t *testing.T) {
			if got := bstFromPostorder(tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bstFromPostorder() = %v, want %v", got, tt.want)
			}
		})
	}

}
