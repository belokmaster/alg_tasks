package main

func minOperations(nums []int) int {
	ans := 0
	n := len(nums)

	for i := 1; i < n-1; i++ {
		if nums[i-1] == 0 {
			nums[i-1] = 1
			nums[i] ^= 1
			nums[i+1] ^= 1
			ans++
		}
	}

	if nums[n-1]+nums[n-2] != 2 {
		return -1
	}

	return ans
}
