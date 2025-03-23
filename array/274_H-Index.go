package main

import "sort"

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	h := 0

	for i, cit := range citations {
		if cit >= i+1 {
			h = i + 1
		} else {
			break
		}
	}

	return h
}
