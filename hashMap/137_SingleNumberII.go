package main

func singleNumber2(nums []int) int {
	hashMap := make(map[int]int)

	for _, num := range nums {
		hashMap[num]++
	}

	for key, value := range hashMap {
		if value == 1 {
			return key
		}
	}

	return nums[0]
}
