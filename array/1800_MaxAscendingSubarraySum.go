package main

import "math"

func maxAscendingSum(nums []int) int {
	ans := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			curSum += nums[i]
			ans = int(math.Max(float64(ans), float64(curSum)))
		} else {
			ans = int(math.Max(float64(ans), float64(curSum)))
			curSum = nums[i]
		}
	}

	return ans
}
