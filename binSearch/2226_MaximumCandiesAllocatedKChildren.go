package main

import "slices"

func canAllocate(candies []int, k int, mid int) bool {
	count := 0
	for _, c := range candies {
		count += c / mid
		if count >= k {
			return true
		}
	}
	return false
}

func maximumCandies(candies []int, k int64) int {
	if k == 0 {
		return 0
	}

	left, right := 1, slices.Max(candies)
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		if canAllocate(candies, int(k), mid) {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
