package main

import "sort"

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func findMin1(nums []int) int {
	sort.Ints(nums)
	return (nums[0])
}
