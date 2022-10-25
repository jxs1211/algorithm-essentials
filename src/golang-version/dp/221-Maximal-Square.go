package dp

// Definition: 以 matrix[i][j] 为右下角元素的最大的全为 1 正方形矩阵的边长为 dp[i][j]。
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	ma := map[byte]int{
		48: 0,
		49: 1,
	}
	// initialize column
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// base case，第一行和第一列的正方形边长
	for i := 0; i < n; i++ {
		dp[0][i] = ma[matrix[0][i]]
	}
	for i := 0; i < m; i++ {
		dp[i][0] = ma[matrix[i][0]]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 值为 0 不可能是正方形的右下角
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = 1 + min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1])
		}
	}
	maxLen := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxLen = max(maxLen, dp[i][j])
		}
	}
	return maxLen * maxLen
}
