package main

import "math"

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && boxes[j] == '1' {
				ans[i] += int(math.Abs(float64(j) - float64(i)))
			}
		}
	}
	return ans
}
