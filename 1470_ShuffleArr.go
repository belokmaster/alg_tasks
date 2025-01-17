package main

func shuffle(nums []int, n int) []int {
	ans := make([]int, 2*n)
	for i := 0; i < n; i++ {
		ans[2*i] = nums[i]
		ans[2*i+1] = nums[n+i]
	}

	return ans
}
