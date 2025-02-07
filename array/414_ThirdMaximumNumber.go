package main

import (
	"slices"
	"sort"
)

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	uniqueDigitsMap := make(map[int]bool)
	uniqueDigits := []int{}

	for _, num := range nums {
		if uniqueDigitsMap[num] {
			continue
		} else {
			uniqueDigits = append(uniqueDigits, num)
			uniqueDigitsMap[num] = true
		}
	}

	sort.Ints(uniqueDigits)
	slices.Reverse(uniqueDigits)
	if len(uniqueDigits) == 1 || len(uniqueDigits) == 2 {
		return uniqueDigits[0]
	} else {
		return uniqueDigits[2]
	}
}
