package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	closestSum := math.MaxInt

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			if abs(currentSum-target) < abs(closestSum-target) {
				closestSum = currentSum
			}
			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				return currentSum
			}
		}
	}

	return closestSum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
