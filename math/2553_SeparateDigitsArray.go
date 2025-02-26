package main

import "slices"

func digitSum(num int) []int {
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	slices.Reverse(digits)
	return digits
}

func separateDigits(nums []int) []int {
	ans := []int{}

	for _, num := range nums {
		x := digitSum(num)
		ans = append(ans, x...)
	}

	return ans
}
