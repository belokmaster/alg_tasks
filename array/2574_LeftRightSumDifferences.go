package main

import "math"

func sum2(arr []int) int {
	ans := 0
	for _, num := range arr {
		ans += num
	}
	return ans
}

func leftRightDifference(nums []int) []int {
	leftSum := 0
	rightSum := sum2(nums)

	ans := make([]int, len(nums))
	for i, num := range nums {
		rightSum -= num
		ans[i] = int(math.Abs(float64(rightSum) - float64(leftSum)))
		leftSum += num
	}

	return ans
}
