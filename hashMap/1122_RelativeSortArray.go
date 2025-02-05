package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	missDigits := []int{}

	hashMap := make(map[int]int)
	for _, num := range arr1 {
		hashMap[num]++
	}

	ans := []int{}
	for _, num := range arr2 {
		if count, exists := hashMap[num]; exists {
			for i := 0; i < count; i++ {
				ans = append(ans, num)
				hashMap[num]--
			}
		}
	}

	for key, value := range hashMap {
		if value != 0 {
			for i := 0; i < value; i++ {
				missDigits = append(missDigits, key)
			}
		}
	}

	sort.Ints(missDigits)
	ans = append(ans, missDigits...)
	return ans
}
