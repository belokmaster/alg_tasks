package main

func binSearch(arr []int, target int) bool {
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

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] > target {
			return false
		} else {
			if binSearch(matrix[i], target) {
				return true
			}
		}
	}
	return false
}
