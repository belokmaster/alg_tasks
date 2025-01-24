package main

import "sort"

func getSneakyNumbers(nums []int) []int {
	sort.Ints(nums)
	ans := make([]int, 0)
	arr := make([]bool, len(nums)+1)

	for _, num := range nums {
		if arr[num] == true {
			ans = append(ans, num)
		}

		if len(ans) == 2 {
			return ans
		}

		arr[num] = true
	}

	return ans
}
