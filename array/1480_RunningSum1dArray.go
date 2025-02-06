package main

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	sum := nums[0]
	ans[0] = sum
	for i := 1; i < len(nums); i++ {
		ans[i] = sum + nums[i]
		sum += nums[i]
	}
	return ans
}
