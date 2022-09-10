package algorithm

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "base",
			nums: []int{},
			want: 0,
		},
		{
			name: "duplicate",
			nums: []int{1, 1, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	path := 1
	val := 1
	t.Logf("%d\n", path<<1)
	t.Logf("%d\n", path<<1&val)
}
