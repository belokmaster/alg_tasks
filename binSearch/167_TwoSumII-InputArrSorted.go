package main

func binSearchLeft(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			ans = mid
			left = mid + 1
		}

		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		x := binSearchLeft(numbers, target-numbers[i])
		if x != -1 {
			return []int{i + 1, x + 1}
		}
	}
	return []int{1, 2}
}
