package main

import (
	"slices"
	"sort"
)

func findMaxK(nums []int) int {
	ans := -1
	sort.Ints(nums)

	for _, num := range nums {
		if num > 0 {
			return ans
		}

		x := num * -1
		if slices.Contains(nums, x) {
			if (x) > ans {
				ans = x
			}
		}
	}

	return ans
}
