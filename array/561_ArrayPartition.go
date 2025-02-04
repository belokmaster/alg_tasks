package main

import "sort"

func arrayPairSum(arr []int) int {
	sort.Ints(arr)

	sum := 0
	for i := 0; i < len(arr); i += 2 {
		sum += arr[i]
	}
	return sum
}
