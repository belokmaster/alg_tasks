package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	mapNums1 := make(map[int]bool)
	mapNums2 := make(map[int]bool)
	ans := [][]int{}
	ans1 := []int{}
	ans2 := []int{}

	lenNums1 := 0
	for _, num := range nums1 {
		if mapNums1[num] {
			continue
		}
		mapNums1[num] = true
		nums1[lenNums1] = num
		lenNums1++
	}
	nums1 = nums1[:lenNums1]

	lenNums2 := 0
	for _, num := range nums2 {
		if mapNums2[num] {
			continue
		}
		mapNums2[num] = true
		nums2[lenNums2] = num
		lenNums2++
	}
	nums2 = nums2[:lenNums2]

	for _, num := range nums1 {
		if !mapNums2[num] {
			ans1 = append(ans1, num)
		}
	}

	for _, num := range nums2 {
		if !mapNums1[num] {
			ans2 = append(ans2, num)
		}
	}

	ans = append(ans, ans1)
	ans = append(ans, ans2)

	return ans
}
