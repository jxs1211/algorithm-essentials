package dp

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Solution2 ![https://labuladong.github.io/algo/3/26/79/]

func lengthOfLIS2(nums []int) int {
	top := make([]int, len(nums))
	// 牌堆数初始化为 0
	piles := 0
	for i := 0; i < len(nums); i++ {
		// 要处理的扑克牌
		pocker := nums[i]
		/***** 搜索左侧边界的二分查找 *****/
		l, r := 0, piles
		for l < r {
			mid := (l + r) >> 1
			if pocker <= top[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		}
		/*********************************/
		// 没找到合适的牌堆，新建一堆
		if l == piles {
			piles++
		}
		// 把这张牌放到牌堆顶
		top[l] = pocker
	}
	return piles
}
