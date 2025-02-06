package main

import "sort"

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))

	for i, num := range nums {
		ans[i] = num * num
	}

	sort.Ints(ans)
	return ans
}
