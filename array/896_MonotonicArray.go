package main

func isLower(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if !(nums[i] <= nums[i+1]) {
			return false
		}
	}
	return true
}

func isHigher(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if !(nums[i] >= nums[i+1]) {
			return false
		}
	}
	return true
}

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	return isHigher(nums) || isLower(nums)
}
