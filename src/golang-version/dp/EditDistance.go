package dp

// Solution1 from top to bottom
// dp[i - 1][j - 1]  replace
// dp[i][j - 1]  insert
// dp[i - 1][j]  delete
// dp return the minimum numbers of opertaions of converting word1 to word2
func dp(word1 string, i int, word2 string, j int, memo *[][]int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}
	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}
	if word1[i] == word2[j] {
		(*memo)[i][j] = dp(word1, i-1, word2, j-1, memo)
		return (*memo)[i][j]
	}
	(*memo)[i][j] = min(
		dp(word1, i-1, word2, j-1, memo)+1,
		min(dp(word1, i, word2, j-1, memo)+1, dp(word1, i-1, word2, j, memo)+1),
	)
	return (*memo)[i][j]
}

func minDistance(word1 string, word2 string) int {
	w1, w2 := len(word1)-1, len(word2)-1
	memo := make([][]int, w1+1)
	for i := 0; i <= w1; i++ {
		s := make([]int, w2+1)
		for j := 0; j <= w2; j++ {
			s[j] = -1
		}
		memo[i] = s
	}
	return dp(word1, w1, word2, w2, &memo)
}
