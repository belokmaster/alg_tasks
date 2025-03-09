package main

import "math"

func minElement(nums []int) int {
	minElement := int(math.MaxInt64)
	ansIndex := 0
	for i, num := range nums {
		if num < minElement {
			ansIndex = i
			minElement = num
		}
	}
	return ansIndex
}

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minIndex := minElement(nums)
		nums[minIndex] *= multiplier
	}
	return nums
}
	