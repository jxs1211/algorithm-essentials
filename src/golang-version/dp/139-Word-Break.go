package dp

// 让你判断 s 是否能被分解成 wordDict 中的单词，
// 反过来想就是判断 wordDict 中的单词是否能拼出 s，那么暴力穷举的思路就是：
func dpWordBreak(s string, i int, wordDict []string, memo *[]int) bool {
	if i == len(s) {
		return true
	} // base case，整个 s 都被拼出来了
	if (*memo)[i] != -1 { // 防止冗余计算
		return (*memo)[i] == 1
	}
	// 遍历所有单词，尝试匹配 s[i..] 的前缀
	for _, word := range wordDict {
		l := len(word)
		if i+l > len(s) {
			continue
		}
		subStr := s[i : i+l]
		if subStr != word {
			continue
		}
		// s[i..] 的前缀被匹配，去尝试匹配 s[i+len..]
		if dpWordBreak(s, i+l, wordDict, memo) {
			(*memo)[i] = 1 // s[i..] 可以被拼出，将结果记入备忘录
			return true
		}
	}
	// s[i..] 不能被拼出，结果记入备忘录
	(*memo)[i] = 0
	return false
}

func wordBreak(s string, wordDict []string) bool {
	memo := make([]int, len(s))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return dpWordBreak(s, 0, wordDict, &memo)
}
