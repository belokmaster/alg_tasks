package main

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestNumber(nums []int) int {
	ans := 0
	answerDist := math.MaxInt

	for _, num := range nums {
		dist := Abs(num)
		if dist < answerDist {
			answerDist = dist
			ans = num
		} else if dist == answerDist && num > ans {
			ans = num
		}
	}

	return ans
}
