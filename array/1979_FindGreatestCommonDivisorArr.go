package main

import "sort"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findGCD(nums []int) int {
	sort.Ints(nums)
	return gcd(nums[0], nums[len(nums)-1])
}
