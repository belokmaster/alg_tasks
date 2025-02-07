package main

func sum1(arr []int) int {
	ans := 0
	for _, num := range arr {
		ans += num
	}
	return ans
}

func findMiddleIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	ans := -1
	rightSum := 0
	leftSum := sum1(nums)

	for i := 0; i < len(nums); i++ {
		if leftSum-nums[i] == rightSum {
			return i
		}
		leftSum -= nums[i]
		rightSum += nums[i]
	}

	return ans
}
