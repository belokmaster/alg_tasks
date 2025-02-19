package main

import (
	"math"
	"sort"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)
	maxDiff := 0

	for i := 0; i < len(nums)-1; i++ {
		a := nums[i]
		b := nums[i+1]
		diff := int(math.Abs(float64(a) - float64(b)))

		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}
