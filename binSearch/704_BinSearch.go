package main

func searchBin(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		// Если элемент в середине больше искомого, ищем в левой половине
		if nums[mid] > target {
			right = mid - 1
		} else {
			// Если элемент в середине меньше искомого, ищем в правой половине
			left = mid + 1
		}
	}

	return -1
}
