package main

// бин поиск индекса последнего положительного
func searchFirstNegative(arr []int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] >= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func countNegatives(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		firstNegativeIndex := searchFirstNegative(grid[i])
		ans += len(grid[i]) - firstNegativeIndex
	}
	return ans
}
