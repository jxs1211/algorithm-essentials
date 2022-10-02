package algorithm

import "testing"

func TestBitAlgorithm(t *testing.T) {
	// https://mp.weixin.qq.com/s/8lJNdnJ0tWm2CapiW_u7XA
	t.Log(('d' ^ ' '))
	t.Log(('d' & '_'))
	t.Log('a' | ' ')
	t.Log('A' | ' ')
}

func Test_pseudoPalindromicPaths(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths(tt.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traversePath(t *testing.T) {
	tests := []struct {
		name  string
		root  *TreeNode
		count *int
		res   *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traversePath(tt.root, tt.count, tt.res)
		})
	}
}

func Test_pseudoPalindromicPaths2(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pseudoPalindromicPaths2(tt.args.root); got != tt.want {
				t.Errorf("pseudoPalindromicPaths2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traversePath2(t *testing.T) {
	type args struct {
		root  *TreeNode
		count *[]int
		res   *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traversePath2(tt.args.root, tt.args.count, tt.args.res)
		})
	}
}
