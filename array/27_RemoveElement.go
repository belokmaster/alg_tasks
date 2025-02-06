package main

func removeElement(nums []int, val int) int {
	ind := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ind] = nums[i]
			ind++
		}
	}
	return ind
}
