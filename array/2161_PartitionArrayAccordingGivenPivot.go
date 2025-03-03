package main

func pivotArray(nums []int, pivot int) []int {
	beforePivot := []int{}
	betweenPivot := []int{}
	afterPivot := []int{}

	for _, num := range nums {
		if num < pivot {
			beforePivot = append(beforePivot, num)
		} else if num == pivot {
			betweenPivot = append(betweenPivot, num)
		} else {
			afterPivot = append(afterPivot, num)
		}
	}

	beforePivot = append(beforePivot, betweenPivot...)
	beforePivot = append(beforePivot, afterPivot...)
	return beforePivot
}
