package main

func sortArrayByParity(nums []int) []int {
	oddIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			temp := nums[oddIndex]
			nums[oddIndex] = nums[i]
			nums[i] = temp
			oddIndex++
		}
	}
	return nums
}
