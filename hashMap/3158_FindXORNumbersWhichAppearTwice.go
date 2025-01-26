package main

func doubleNumber(nums []int) []int {
	hashMap := make(map[int]int)
	ans := []int{}

	for _, num := range nums {
		hashMap[num]++
	}

	for key, value := range hashMap {
		if value == 2 {
			ans = append(ans, key)
		}
	}

	return ans
}

func duplicateNumbersXOR(nums []int) int {
	digits := doubleNumber(nums)
	result := 0

	for _, num := range digits {
		result ^= num
	}

	return result
}
