package main

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	hashTable := make(map[int][]int)
	for i, num := range nums {
		hashTable[num] = append(hashTable[num], i)
	}

	for _, value := range hashTable {
		if len(value) > 1 {
			for i := 0; i < len(value)-1; i++ {
				if int(math.Abs(float64(value[i])-float64(value[i+1]))) <= k {
					return true
				}
			}
		}
	}

	return false
}
