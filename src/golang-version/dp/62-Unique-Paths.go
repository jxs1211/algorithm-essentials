package dp

// dp[i][j] means the sum of paths can walk through from original point to the destincation(dp[i][j])
func dpUniquePaths(i, j int, memo *[][]int) int {
	if i == 0 && j == 0 {
		return 1
	}
	if i < 0 || j < 0 {
		return 0
	}
	if (*memo)[i][j] > 0 {
		return (*memo)[i][j]
	}
	(*memo)[i][j] = dpUniquePaths(i-1, j, memo) + dpUniquePaths(i, j-1, memo)
	return (*memo)[i][j]
}
func uniquePaths(m int, n int) int {
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	return dpUniquePaths(m-1, n-1, &memo)
}
