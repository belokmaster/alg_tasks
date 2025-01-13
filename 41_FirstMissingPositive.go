package main

import "sort"

func firstMissingPositive(nums []int) int {
	minDigit := 1
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			if nums[i] == minDigit {
				minDigit++
			}
		}
	}

	return minDigit
}
