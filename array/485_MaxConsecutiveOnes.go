package main

func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > ans {
				ans = count
			}
		} else {
			count = 0
		}
	}

	return ans
}
