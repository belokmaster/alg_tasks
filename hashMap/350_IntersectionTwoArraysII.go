package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	hashMap := make(map[int]int)
	ans := make([]int, 0)

	for _, num := range nums1 {
		hashMap[num]++
	}

	for _, num := range nums2 {
		if count, exists := hashMap[num]; exists && count > 0 {
			ans = append(ans, num)
			hashMap[num]--
		}
	}

	return ans
}
