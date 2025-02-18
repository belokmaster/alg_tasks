package main

func sortColors(nums []int) {
	colorsMap := make(map[int]int)

	for _, color := range nums {
		colorsMap[color]++
	}

	index := 0
	for color := 0; color < 3; color++ {
		count := colorsMap[color]
		for j := 0; j < count; j++ {
			nums[index] = color
			index++
		}
	}
}
