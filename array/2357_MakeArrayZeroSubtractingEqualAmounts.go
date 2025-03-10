package main

import "sort"

func minusDigit(nums []int, k int) []int {
	ans := make([]int, len(nums))
	copy(ans, nums) // копирование только если len равен
	for i := range ans {
		ans[i] -= k
	}
	return ans
}

func removeZeros(nums []int) []int {
	result := []int{}
	for i := range nums {
		if nums[i] > 0 {
			result = nums[i:]
			break
		}
	}
	return result
}

func minimumOperations(nums []int) int {
	count := 0

	for len(nums) != 0 {
		sort.Ints(nums)
		nums = removeZeros(nums)
		if len(nums) == 0 {
			break
		}

		k := nums[0]
		if k < 0 {
			break
		}

		nums = minusDigit(nums, k)
		count++
	}

	return count
}
