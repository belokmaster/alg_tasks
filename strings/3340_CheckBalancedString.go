package main

import "strconv"

func isBalanced(num string) bool {
	evenSum := 0
	oddSum := 0

	for i := 0; i < len(num); i++ {
		x, _ := strconv.Atoi(string(num[i]))
		if i%2 == 0 {
			evenSum += x
		} else {
			oddSum += x
		}
	}

	return evenSum == oddSum
}
