package main

func BinaryFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func BinaryLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func searchRange(nums []int, target int) []int {
	answer := make([]int, 2)
	answer[0] = BinaryFirst(nums, target)
	answer[1] = BinaryLast(nums, target)
	return answer
}
