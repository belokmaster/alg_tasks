package main

import "sort"

func BinaryFirst1(nums []int, key int) int {
	left := 0
	right := len(nums) - 1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < key {
			left = mid + 1
		} else if nums[mid] == key {
			ans = mid
			right = mid - 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func BinaryLast1(nums []int, key int) int {
	left := 0
	right := len(nums) - 1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < key {
			left = mid + 1
		} else if nums[mid] == key {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	first := BinaryFirst1(nums, target)
	last := BinaryLast1(nums, target)

	if first == -1 {
		return []int{}
	}

	ans := []int{}
	for i := first; i <= last; i++ {
		ans = append(ans, i)
	}

	return ans
}
