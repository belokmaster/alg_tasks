package main

import "fmt"

func checkRow(arr []int) bool {
	n := len(arr)
	sum := (n * (n + 1)) / 2
	mapDigit := make(map[int]bool)
	for _, num := range arr {
		sum -= num
		if mapDigit[num] {
			return false
		}
		mapDigit[num] = true
	}
	return sum == 0
}

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		sum := (n * (n + 1)) / 2
		fmt.Println(sum)
		if !(checkRow(matrix[i])) {
			return false
		}
		fmt.Println(sum)
		for j := 0; j < n; j++ {
			sum -= matrix[j][i]
		}

		if sum != 0 {
			return false
		}
	}

	return true
}
