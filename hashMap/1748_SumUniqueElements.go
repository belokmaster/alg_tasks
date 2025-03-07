package main

func sumOfUnique(nums []int) int {
	uniqueDigits := make(map[int]int)
	ans := 0

	for _, num := range nums {
		uniqueDigits[num]++
	}

	for key, value := range uniqueDigits {
		if value == 1 {
			ans += key
		}
	}

	return ans
}
