package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	ans := 0

	for i := 0; i < len(costs); i++ {
		if costs[i] <= coins {
			coins -= costs[i]
			ans++
		}
	}
	return ans
}
