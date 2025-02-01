package main

import "sort"

func uniqueOccurrences(arr []int) bool {
	hashMap := make(map[int]int)
	sort.Ints(arr)
	for _, num := range arr {
		hashMap[num]++
	}

	boolDigits := make(map[int]bool)
	for _, value := range hashMap {
		if boolDigits[value] {
			return false
		}
		boolDigits[value] = true
	}
	return true
}
