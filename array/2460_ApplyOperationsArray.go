package main

func applyOperations(nums []int) []int {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	ans := []int{}
	count := 0
	for _, num := range nums {
		if num != 0 {
			ans = append(ans, num)
			count++
		}
	}

	for i := 0; i < n-count; i++ {
		ans = append(ans, 0)
	}

	return ans
}
