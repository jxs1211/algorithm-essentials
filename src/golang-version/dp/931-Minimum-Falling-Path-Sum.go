package dp

import "math"

// dp means the minimize sum of the path falling from first line to the position of matrix[i][j]
func dpFalling(matrix [][]int, i, j int, memo *[][]int) int {
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
		return 99999
	}
	// base case
	if i == 0 {
		return matrix[0][j]
	}
	if (*memo)[i][j] != 66666 {
		return (*memo)[i][j]
	}

	(*memo)[i][j] = matrix[i][j] + min(dpFalling(matrix, i-1, j-1, memo), min(dpFalling(matrix, i-1, j, memo), dpFalling(matrix, i-1, j+1, memo)))
	return (*memo)[i][j]
}

func minFallingPathSum(matrix [][]int) int {
	l := len(matrix)
	memo := make([][]int, l)
	for i := 0; i < l; i++ {
		s := make([]int, l)
		for j := 0; j < l; j++ {
			s[j] = 66666
		}
		memo[i] = s
	}
	res := math.MaxInt
	for i := 0; i < l; i++ {
		res = min(res, dpFalling(matrix, l-1, i, &memo))
	}
	return res
}
