package main

import "slices"

func alternateDigitSum(num int) int {
	sum := 0
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	slices.Reverse(digits)
	for i, digit := range digits {
		if i%2 == 0 {
			sum += digit
		} else {
			sum -= digit
		}
	}

	return sum
}
