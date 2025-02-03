package main

func removeDuplicates(nums []int) int {
	prevInt := 0
	indPos := 1
	for i := 1; i < len(nums); i++ {
		if nums[prevInt] != nums[i] {
			nums[indPos] = nums[i]
			indPos++
			prevInt++
		}
	}
	return indPos
}
