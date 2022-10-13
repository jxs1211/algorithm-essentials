package algorithm

// postorder
// total possible results = possible result of left sub tree * possible result of right sub tree
func count(lo, hi int, cache [][]int) int {
	if lo > hi {
		return 1
	}
	if cache[lo][hi] != 0 {
		return cache[lo][hi]
	}
	res := 0
	for i := lo; i <= hi; i++ {
		l := count(lo, i-1, cache)
		r := count(i+1, hi, cache)
		res += l * r
	}
	cache[lo][hi] = res
	return res
}

func getCache(n int) [][]int {
	s := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		s[i] = make([]int, n+1)
	}
	return s
}
func numTrees(n int) int {
	cache := getCache(n)
	return count(1, n, cache)
}
