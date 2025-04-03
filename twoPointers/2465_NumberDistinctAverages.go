package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := make(map[float64]bool)

	for i := 0; i < n/2; i++ {
		x := float64(nums[i]+nums[n-i-1]) / 2
		ans[x] = true
	}

	return len(ans)
}
