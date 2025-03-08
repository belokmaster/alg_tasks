package main

import "sort"

func numberGame(nums []int) []int {
	ans := []int{}
	n := len(nums) / 2
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		x := nums[0]
		y := nums[1]
		nums = nums[2:]
		ans = append(ans, y)
		ans = append(ans, x)
	}

	return ans
}
