package main

func findFirstNonNegative(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func maximumCount(nums []int) int {
	n := len(nums)
	indFirstPos := findFirstNonNegative(nums)
	negCount := indFirstPos
	posCount := n - indFirstPos
	zeroCount := 0

	for i := indFirstPos; i < n; i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			break
		}
	}

	posCount -= zeroCount
	if posCount > negCount {
		return posCount
	}
	return negCount
}
