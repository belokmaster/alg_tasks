package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	//максимальная сумма, которую можно украсть до текущего дома
	prevMax := 0
	// максимальная сумма, которую можно украсть до текущего дома, включая этот дом.
	currMax := 0

	for _, num := range nums {
		temp := currMax
		currMax = max(currMax, prevMax+num)
		prevMax = temp
	}

	return currMax
}
