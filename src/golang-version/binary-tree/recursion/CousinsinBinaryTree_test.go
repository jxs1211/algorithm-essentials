package algorithm

import (
	"fmt"
	"testing"
)

func Test_isCousins(t *testing.T) {
	// [1,2,3,null,4,null,5]
	// 5
	// 4
	// 	[1,null,2,3,null,null,4,null,5]
	// 1
	// 3
	tests := []struct {
		name string
		root *TreeNode
		x    int
		y    int
		want bool
	}{
		{
			name: "base",
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:   4,
							Right: &TreeNode{Val: 5}},
					},
				},
			},
			x:    1,
			y:    3,
			want: false,
		},
		{
			name: "cousins nodes existed",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 5},
				},
			},
			x:    5,
			y:    4,
			want: true,
		},
		{
			name: "the parents is the same one",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			x:    2,
			y:    3,
			want: false,
		},
		{
			name: "the parents is the same one",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			x:    2,
			y:    3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCousins(tt.root, tt.x, tt.y); got != tt.want {
				t.Errorf("isCousins() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_traverse121(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		root        *TreeNode
// 		parent      *TreeNode
// 		nodeXParent *TreeNode
// 		nodeYParent *TreeNode
// 		level       int
// 		nodeXDepth  *int
// 		nodeYDepth  *int
// 		x           int
// 		y           int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			traverse121(tt.root, tt.parent, tt.nodeXParent, tt.nodeYParent, tt.level, tt.nodeXDepth, tt.nodeYDepth, tt.x, tt.y)
// 		})
// 	}
// }

func modifyPointer(parent *TreeNode) *TreeNode {
	println("before: ", parent)
	parent = &TreeNode{Val: 1}
	println("after: ", parent)
	return parent
}

func TestPassStructAsPoint(t *testing.T) {
	var p *TreeNode
	println(p)
	p2 := modifyPointer(p)
	fmt.Println(p)
	fmt.Println(p2)
}
