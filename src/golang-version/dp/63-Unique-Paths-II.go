package dp

// dp[i][j] means the sum of paths can walk through from original point to the destincation(dp[i][j])

func dp2(grid [][]int, m, n int, memo *[][]int) int {
	// crossover or obstacle
	if m < 0 || m >= len(grid) || n < 0 || n >= len(grid[0]) || grid[m][n] == 1 {
		return 0
	}
	if m == 0 && n == 0 {
		return 1
	}
	if (*memo)[m][n] > 0 {
		return (*memo)[m][n]
	}
	(*memo)[m][n] = dp2(grid, m-1, n, memo) + dp2(grid, m, n-1, memo)
	return (*memo)[m][n]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	return dp2(obstacleGrid, m-1, n-1, &memo)
}
