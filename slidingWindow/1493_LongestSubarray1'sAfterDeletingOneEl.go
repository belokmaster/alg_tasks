package main

import "math"

func longestSubarray(nums []int) int {
	ans := 0
	left := 0
	zeroCount := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		ans = int(math.Max(float64(ans), float64(right-left)))
	}

	return ans
}
