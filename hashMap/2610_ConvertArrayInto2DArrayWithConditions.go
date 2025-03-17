package main

func findMatrix(nums []int) [][]int {
	digits := make(map[int]int)
	ans := [][]int{}

	maxCount := 0
	for _, num := range nums {
		digits[num]++
		if digits[num] > maxCount {
			maxCount = digits[num]
		}
	}

	for i := 0; i < maxCount; i++ {
		arr := []int{}
		for key, num := range digits {
			if num > 0 {
				digits[key]--
				arr = append(arr, key)
			}
		}

		if len(arr) > 0 {
			ans = append(ans, arr)
		}
	}

	return ans
}
