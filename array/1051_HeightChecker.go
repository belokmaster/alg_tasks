package main

import "sort"

func heightChecker(heights []int) int {
	expected := append([]int{}, heights...)
	sort.Ints(expected)

	ans := 0
	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			ans++
		}
	}
	return ans
}
