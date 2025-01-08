package main

func peakIndexInMountainArray(arr []int) int {
	// Первый и последний элементы не могут быть пиком в горном массиве
	left := 1
	right := len(arr) - 2

	for right > left {
		mid := left + (right-left)/2

		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
