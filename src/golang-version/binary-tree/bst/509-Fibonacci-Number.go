package algorithm

// Solution1
func fib(n int) int {
	if n < 2 {
		return n
	}
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

// Solution2
// func fib(n int) int {
//     if n < 2 {return n}
//     dp := make([]int, n + 1)
//     dp[0], dp[1] = 0, 1
//     for i := 2; i <= n; i++ {
//         dp[i] = dp[i - 1] + dp[i - 2]
//     }
//     return dp[n]
// }

// Solution3
// func doFib(n int, memo map[int]int) int {
//     if n < 2 {return n}
//     if v, ok := memo[n]; ok {return v}
//     memo[n] = doFib(n - 1, memo) + doFib(n - 2, memo)
//     return memo[n]
// }
// func fib(n int) int {
//     memo := map[int]int{}
//     return doFib(n, memo)
// }

// Solution4
// func fib(n int) int {
//     if n < 2 {return n}
//     return fib(n - 1) + fib(n - 2)
// }
