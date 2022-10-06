package algorithm

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_minCameraCover(t *testing.T) {
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
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_traverse55(t *testing.T) {
	type args struct {
		root      *TreeNode
		hasParent bool
		res       *int
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
			if got := traverse55(tt.args.root, tt.args.hasParent, tt.args.res); got != tt.want {
				t.Errorf("traverse55() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	fmt.Println([]byte{'a' + 0, 'a' + 1})
	a := []byte{'a' + 0, 'a' + 1}
	a = append(a, 'a'+3)
	fmt.Println(string(a))
	tests := []struct {
		name string
		b    []byte
		want []byte
	}{
		{
			name: "base",
			b:    []byte{'a', 'b', 'c'},
			want: []byte{'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse2(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want []byte
	}{
		{
			name: "base",
			b:    []byte{'a', 'b', 'c'},
			want: []byte{'c', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse2(tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse2() = %v, want %v", got, tt.want)
			}
		})
	}
}
