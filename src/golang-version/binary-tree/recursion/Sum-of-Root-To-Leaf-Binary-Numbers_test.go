package algorithm

import (
	"fmt"
	"testing"
)

func Test_sumRootToLeaf(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "base",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRootToLeaf(tt.root); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfsSumRootToLeaf(t *testing.T) {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)
	// 0010 | 1010
	// type args struct {
	// 	root *TreeNode
	// 	res  *int
	// 	path *int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		dfsSumRootToLeaf(tt.args.root, tt.args.res, tt.args.path)
	// 	})
	// }
}
