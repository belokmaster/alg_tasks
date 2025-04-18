package main

import "math"

func countKDifference(nums []int, k int) int {
	ans := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) == k {
				ans++
			}
		}
	}

	return ans
}
