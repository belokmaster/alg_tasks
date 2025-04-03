package main

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	ans := math.MaxFloat64

	for i := 0; i < n/2; i++ {
		x := (float64(nums[i]) + float64(nums[n-i-1])) / 2
		if x < ans {
			ans = x
		}
	}

	return ans
}
