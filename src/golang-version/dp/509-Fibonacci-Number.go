package dp

// Solution1 best one
// from bottom to top
// TC ON, SC O1
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

// from bottom to top
// TC ON, SC ON
func fib2(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Solution2
// from top to bottom
// TC ON, SC ON
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
// from top to bottom with memo
// TC ON, SC ON
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
// from top to bottom without memo
// TC O2^N, SC ON
// func fib(n int) int {
//     if n < 2 {return n}
//     return fib(n - 1) + fib(n - 2)
// }
