package main

import (
	"fmt"
	"sort"
)

func checkFullSorted(nums, sortNums []int) bool {
	for i := 0; i < len(nums); i++ {
		if sortNums[i] != nums[i] {
			return false
		}
	}
	return true
}

// не забывать что слайс передается по ссылке. надо копировать, чтобы избегать ошибок
func check(nums []int) bool {
	newNums := nums
	sortNums := make([]int, len(nums))
	copy(sortNums, nums)
	sort.Ints(sortNums)

	if checkFullSorted(newNums, sortNums) {
		fmt.Println(newNums)
		return true
	}
	n := len(newNums) - 1
	i := 0
	for i < n {
		newNums = append(newNums, newNums[0])
		newNums = newNums[1:]

		if checkFullSorted(newNums, sortNums) {
			return true
		}

		if newNums[n] < newNums[n-1] {
			return false
		}
		i++
	}
	return true
}
