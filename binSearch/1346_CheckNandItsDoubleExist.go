package main

import "sort"

func binSearch5(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			target := arr[i] / 2
			ind := binSearch5(arr, target)
			if ind != -1 && ind != i {
				return true
			}
		}
	}
	return false
}
