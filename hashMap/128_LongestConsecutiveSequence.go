package main

import (
	"math"
	"sort"
)

func longestConsecutive(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}

	ans := 0
	arr := []int{}
	for key := range hashMap {
		arr = append(arr, key)
	}

	if len(arr) == 1 {
		return 1
	}

	sort.Ints(arr)
	count := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]+1 {
			ans = int(math.Max(float64(ans), float64(count)))
			count = 1
		} else {
			count++
			ans = int(math.Max(float64(ans), float64(count)))
		}
	}

	return ans
}
