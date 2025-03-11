package main

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	hashMap1 := make(map[int]int)
	hashMap2 := make(map[int]int)

	for _, num := range nums1 {
		hashMap1[num]++
	}

	for _, num := range nums2 {
		hashMap2[num]++
	}

	ans := make([]int, 2)
	for k := range hashMap2 {
		ans[0] += hashMap1[k]
	}

	for k := range hashMap1 {
		ans[1] += hashMap2[k]
	}

	return ans
}
