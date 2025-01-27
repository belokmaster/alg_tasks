package main

func binSearch2(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		}

		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if binSearch2(matrix[i], target) {
			return true
		}
	}
	return false
}
