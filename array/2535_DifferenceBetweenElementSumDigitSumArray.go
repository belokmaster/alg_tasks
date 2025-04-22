package main

import "math"

func differenceOfSum(nums []int) int {
	sumElements := 0
	sumDigits := 0
	for _, num := range nums {
		sumElements += num
		for num > 0 {
			sumDigits += num % 10
			num /= 10
		}
	}
	return int(math.Abs(float64(sumDigits - sumElements)))
}
