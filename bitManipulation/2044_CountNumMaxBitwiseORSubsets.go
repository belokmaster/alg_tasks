package main

func countMaxOrSubsets(nums []int) int {
	maxOr := 0
	subsets := []int{0}

	for _, num := range nums {
		newSubsets := []int{}
		for _, val := range subsets {
			newSubsets = append(newSubsets, val|num)
		}
		subsets = append(subsets, newSubsets...)
	}

	for _, val := range subsets {
		maxOr = max(maxOr, val)
	}

	ans := 0
	for _, val := range subsets {
		if val == maxOr {
			ans++
		}
	}

	return ans
}
