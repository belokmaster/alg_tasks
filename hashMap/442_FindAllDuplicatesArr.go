package main

import "sort"

// через булеву мапу, отмечающую повторки.
func findDuplicates(nums []int) []int {
	arr := make([]bool, len(nums)+1)
	ans := []int{}

	for _, num := range nums {
		if arr[num] {
			ans = append(ans, num)
		}
		arr[num] = true
	}

	return ans
}

// через мапу интов, где отмечается кол-во вхождений. брут форс
func findDuplicates1(nums []int) []int {
	hashMap := make(map[int]int)
	sort.Ints(nums)
	for _, num := range nums {
		hashMap[num]++
	}

	ans := []int{}
	for key, value := range hashMap {
		if value == 2 {
			ans = append(ans, key)
		}
	}

	return ans
}
