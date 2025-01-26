package main

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)

	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			i += 2
		} else {
			return nums[i]
		}
	}

	return nums[i]
}
