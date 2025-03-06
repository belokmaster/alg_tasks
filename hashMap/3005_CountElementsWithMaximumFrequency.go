package main

func maxFrequencyElements(nums []int) int {
	hashMap := make(map[int]int)
	maxFreq := 0
	count := 0

	for _, num := range nums {
		hashMap[num]++
		if hashMap[num] > maxFreq {
			maxFreq = hashMap[num]
		}
	}

	for _, freq := range hashMap {
		if freq == maxFreq {
			count += freq
		}
	}

	return count
}
