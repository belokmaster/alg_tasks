package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	ans := [][]int{}
	minDiff := 2147483647

	for i := 0; i < len(arr)-1; i++ {
		a := arr[i]
		b := arr[i+1]
		diff := int(math.Abs(float64(a) - float64(b)))
		if diff < minDiff {
			minDiff = diff
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		a := arr[i]
		b := arr[i+1]
		diff := int(math.Abs(float64(a) - float64(b)))
		if diff == minDiff {
			ans = append(ans, []int{a, b})
		}
	}

	return ans
}
