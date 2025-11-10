package main

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	maxCur := nums[0]
	maxAns := nums[0]

	for i := 1; i < len(nums); i++ {
		maxCur = max1(nums[i], nums[i]+maxCur)
		maxAns = max1(maxAns, maxCur)
	}

	return maxAns
}
