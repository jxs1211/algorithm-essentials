package algorithm

// bit algorithm 原理见 https://mp.weixin.qq.com/s/8lJNdnJ0tWm2CapiW_u7XA
// solution1
func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	res := 0
	traversePath(root, &count, &res)
	return res
}

func traversePath(root *TreeNode, count *int, res *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		// 遇到叶子节点，判断路径是否是伪回文串
		*count = *count ^ (1 << root.Val)
		// 判断二进制中只有一位 1，原理见 https://mp.weixin.qq.com/s/8lJNdnJ0tWm2CapiW_u7XA
		if *count&(*count-1) == 0 {
			*res++
		}
		*count = *count ^ (1 << root.Val)
		return
	}
	*count = *count ^ (1 << root.Val)
	traversePath(root.Left, count, res)
	traversePath(root.Right, count, res)
	*count = *count ^ (1 << root.Val)
}

// solution2
func pseudoPalindromicPaths2(root *TreeNode) int {
	count := make([]int, 10)
	res := 0
	traversePath2(root, &count, &res)
	return res
}

func traversePath2(root *TreeNode, count *[]int, res *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		(*count)[root.Val]++
		odd := 0
		for _, v := range *count {
			if v%2 == 1 {
				odd++
			}
		}
		if odd <= 1 {
			*res++
		}
		(*count)[root.Val]--
		return
	}
	(*count)[root.Val]++
	traversePath2(root.Left, count, res)
	traversePath2(root.Right, count, res)
	(*count)[root.Val]--
}
