package main

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if i != j {
				if nums[j] < nums[i] {
					count++
				}
			}
		}
		ans[i] = count
	}

	return ans
}
