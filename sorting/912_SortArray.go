package main

// обычный mergesort
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	var newArr []int
	l := 0
	r := 0

	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			newArr = append(newArr, right[r])
			r++
		} else {
			newArr = append(newArr, left[l])
			l++
		}
	}

	if l < len(left) {
		newArr = append(newArr, left[l:]...)
	}
	if r < len(right) {
		newArr = append(newArr, right[r:]...)
	}

	return newArr
}
