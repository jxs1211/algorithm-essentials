package algorithm

import (
	"math"
	"sort"
)

// solution1

func twoSum(nums []int, target int) []int {
	rec := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if idx, ok := rec[diff]; ok {
			return []int{i, idx}
		}
		rec[v] = i
	}
	return []int{}
}

// solution2

func twoSum2(nums []int, target int) []int {
	cp := copySlice(nums)
	mySort(cp)

	left, right := 0, len(nums)-1
	for left <= right {
		sum := cp[left] + cp[right]
		if sum == target {
			return findOriginIndex(nums, cp, left, right)
			// return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func copySlice(src []int) []int {
	cp := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		cp[i] = src[i]
	}
	return cp
}

func findOriginIndex(nums []int, cp []int, left, right int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if cp[left] == nums[i] {
			res[0] = i
			nums[i] = math.MaxInt
			break
		}
	}
	for i := 0; i < len(nums); i++ {
		if cp[right] == nums[i] {
			res[1] = i
			break
		}
	}
	return res
}

func mySort(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}
