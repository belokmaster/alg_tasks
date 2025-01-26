package main

func singleNumber3(nums []int) []int {
	hashMap := make(map[int]int)
	ans := []int{}

	for _, num := range nums {
		hashMap[num]++
	}

	for key, value := range hashMap {
		if value == 1 {
			ans = append(ans, key)
		}
		if len(ans) == 2 {
			return ans
		}
	}

	return ans
}
