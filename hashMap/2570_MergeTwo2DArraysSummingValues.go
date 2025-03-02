package main

import "sort"

type kv struct {
	Key   int
	Value int
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	hashMap := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		id, val := nums1[i][0], nums1[i][1]
		hashMap[id] = val
	}

	for i := 0; i < len(nums2); i++ {
		id, val := nums2[i][0], nums2[i][1]
		if value, exists := hashMap[id]; exists {
			hashMap[id] = value + val
		} else {
			hashMap[id] = val
		}
	}

	var sortedSlice []kv
	for key, val := range hashMap {
		sortedSlice = append(sortedSlice, kv{key, val})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Key < sortedSlice[j].Key
	})

	ans := [][]int{}
	for i := 0; i < len(sortedSlice); i++ {
		ans = append(ans, []int{sortedSlice[i].Key, sortedSlice[i].Value})
	}

	return ans
}
