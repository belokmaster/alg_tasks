package main

func count(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	hashMap := make(map[int]bool)
	for _, num := range nums1 {
		hashMap[num] = true
	}

	ans := []int{}
	for _, num := range nums2 {
		if hashMap[num] && !count(ans, num) {
			ans = append(ans, num)
		}
	}

	return ans
}
