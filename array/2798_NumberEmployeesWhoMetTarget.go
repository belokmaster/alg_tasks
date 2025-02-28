package main

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	ans := 0

	for _, num := range hours {
		if num >= target {
			ans++
		}
	}

	return ans
}
