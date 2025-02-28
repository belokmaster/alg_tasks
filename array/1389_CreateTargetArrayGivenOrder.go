package main

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, 0, len(nums))

	for i, v := range index {
		ans = append(ans[:v+1], ans[v:]...)
		ans[v] = nums[i]
	}

	return ans
}
