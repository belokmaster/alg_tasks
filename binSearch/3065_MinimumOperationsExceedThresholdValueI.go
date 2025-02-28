package main

import "sort"

func minOperations(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
