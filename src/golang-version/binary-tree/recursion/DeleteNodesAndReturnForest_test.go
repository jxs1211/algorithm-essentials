package algorithm

import (
	"reflect"
	"testing"
)

func Test_delNodes(t *testing.T) {
	type args struct {
		root      *TreeNode
		to_delete []int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delNodes(tt.args.root, tt.args.to_delete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_delete(t *testing.T) {
	type args struct {
		root      *TreeNode
		hasParent bool
		set       map[int]bool
		res       *[]*TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delete(tt.args.root, tt.args.hasParent, tt.args.set, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getMemoMap(n int) map[int]map[int]int {
	memo := make(map[int]map[int]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make(map[int]int, n)
	}
	return memo
}

func TestSomething(t *testing.T) {
	// m := getMemoMap(10)
	// t.Log(m)
	m := make(map[int]int, 10)
	t.Log(m)
	// s := []int{1, 2, 3, 4}
	// fmt.Printf("%p\n", s)
	// s2 := s[1:2]
	// fmt.Printf("%p\n", s2)
	// s2[0] = 10
	// fmt.Printf("%v\n", s)

	// m := map[int]int{}
	// m2 := map[int]map[int]int{
	// 	1: map[int]int{},
	// }

	// s := make([][]int, 10)
	// s := [][]int{}
	// for i := 0; i < 10; i++ {
	// 	s = append(s, make([]int, 10))
	// }
	// t.Log(len(s), len(s[0]))
	// s[0][1] = 1
	// t.Log(s)
	// m2 := make(map[int]map[int]int, 10)
	// m2[1][2] = 3
	// n := 11
	// fmt.Printf("%b, >>2: %b\n", n, n>>2)
	// n |= n >> 2
	// fmt.Printf("%b\n", n) // 11
	// m := 11
	// fmt.Printf("%b, >>2: %b\n", m, m>>2)
	// m &= (m >> 2)
	// fmt.Printf("%b\n", m) // 10
	// f := 11
	// fmt.Printf("%b, >>2: %b\n", f, f>>2)
	// f ^= (f >> 2)         // 1011 0010
	// fmt.Printf("%b\n", f) // 1001
}
