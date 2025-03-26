package main

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	sort.Ints(nums)
	target := nums[len(nums)/2]

	ans := 0
	for _, num := range nums {
		ans += int(math.Abs(float64(num - target)))
	}

	return ans
}
