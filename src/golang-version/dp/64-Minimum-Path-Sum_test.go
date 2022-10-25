package dp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := minPathSum(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum2(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := minPathSum2(tt.args.grid); got != tt.want {
				t.Errorf("minPathSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doMinPathSum(t *testing.T) {
	type args struct {
		grid [][]int
		i    int
		j    int
		memo *[][]int
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
			if got := doMinPathSum(tt.args.grid, tt.args.i, tt.args.j, tt.args.memo); got != tt.want {
				t.Errorf("doMinPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func TestByteToInt(t *testing.T) {
	fmt.Printf("%T, %T, %T\n", 1, int(1), int8(1))
	fmt.Printf("%d\n", BytesToInt([]byte{'1'}))
}
