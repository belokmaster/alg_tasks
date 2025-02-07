package main

func dominantIndex(nums []int) int {
	maxElement := nums[0]
	maxElementIndex := 0
	for i, num := range nums {
		if num > maxElement {
			maxElement = num
			maxElementIndex = i
		}
	}

	for _, num := range nums {
		if num*2 > maxElement && num != maxElement {
			return -1
		}
	}
	return maxElementIndex
}
