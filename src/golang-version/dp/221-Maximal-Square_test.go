package dp

import "testing"

func Test_maximalSquare(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name: "base",
			matrix: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'0', '0', '1', '1', '1'},
			},
			want: 16,
		},
		{
			name: "case2",
			matrix: [][]byte{
				{'0', '1'},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
