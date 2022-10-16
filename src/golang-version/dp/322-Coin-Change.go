package dp

// status: amount
// choice: coins[i]
// the formula of transforming status: dp[i] = min(dp[i], dp[i - coin] + 1)
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1 // initial max value for picking minimum
	}
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
