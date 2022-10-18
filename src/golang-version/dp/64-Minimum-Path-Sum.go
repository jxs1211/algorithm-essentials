package dp

import "math"

// dp[i][j] means the minimum of the sum of the path's value at current postion[i, j]
// from bottom to top
func minPathSum(grid [][]int) int {
	len1, len2 := len(grid), len(grid[0])
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < len1; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len2; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[len1-1][len2-1]
}

// doMinPathSum means the minimum of the sum of the path's value at current postion[i, j] througout from the position of left top
// from bottom to top
func minPathSum2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		column := make([]int, n)
		for j := 0; j < n; j++ {
			column[j] = -1
		}
		memo[i] = column
	}
	return doMinPathSum(grid, m-1, n-1, &memo)
}

func doMinPathSum(grid [][]int, i, j int, memo *[][]int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 { // avoid to be returned by min function
		return math.MaxInt
	}
	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}
	// 左边和上面的最小路径和加上 grid[i][j]
	// 就是到达 (i, j) 的最小路径和
	(*memo)[i][j] = grid[i][j] + min(
		doMinPathSum(grid, i, j-1, memo),
		doMinPathSum(grid, i-1, j, memo))
	return (*memo)[i][j]
}
