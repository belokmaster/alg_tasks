package main

func longestMonotonicSubarray(nums []int) int {
	ans := 1
	posLen := 1
	negLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			posLen++
			negLen = 1
		} else if nums[i] < nums[i-1] {
			negLen++
			posLen = 1
		} else {
			posLen = 1
			negLen = 1
		}

		if posLen > ans {
			ans = posLen
		}
		if negLen > ans {
			ans = negLen
		}
	}

	return ans
}
