package main

func transformArray(nums []int) []int {
	i := 0

	for _, v := range nums {
		if v%2 == 0 {
			nums[i] = 0
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 1
	}

	return nums
}
