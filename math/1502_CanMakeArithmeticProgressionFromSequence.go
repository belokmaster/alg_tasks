package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) == 2 {
		return true
	}

	sort.Ints(arr)
	a1 := arr[0]
	a2 := arr[1]
	d := a2 - a1

	for i := 2; i < len(arr); i++ {
		if arr[i]-i*d != a1 {
			return false
		}
	}

	return true
}
