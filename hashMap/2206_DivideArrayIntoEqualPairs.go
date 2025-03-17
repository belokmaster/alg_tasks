package main

func divideArray(nums []int) bool {
	countElements := make(map[int]int)

	for _, num := range nums {
		countElements[num]++
	}

	for _, value := range countElements {
		if value%2 != 0 {
			return false
		}
	}

	return true
}
