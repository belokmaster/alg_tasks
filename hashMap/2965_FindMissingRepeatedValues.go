package main

import "fmt"

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid[0])
	hashMap := make(map[int]int)
	ans := (1 + n*n) * (n * n) / 2

	for i := range grid {
		for _, j := range grid[i] {
			hashMap[j]++
		}
	}

	answer := []int{}
	for k, v := range hashMap {
		fmt.Println(ans, k, v)
		if v == 2 {
			answer = append(answer, k)
		}
		ans -= k
	}
	answer = append(answer, ans)

	return answer
}
