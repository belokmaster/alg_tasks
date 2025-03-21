package main

import "math"

func max(nums []int) int {
	maxDigit := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxDigit {
			maxDigit = nums[i]
		}
	}
	return maxDigit
}

func divideArr(nums []int, div int) int {
	count := 0

	for _, num := range nums {
		count += int(math.Ceil(float64(num) / float64(div)))
	}

	return count
}

func minEatingSpeed(piles []int, h int) int {
	right := max(piles)
	left := 1
	ans := right

	for left <= right {
		mid := (left + right) / 2
		count := divideArr(piles, mid)

		if count <= h {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
