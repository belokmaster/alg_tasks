package main

import "math"

func divisors(n int) []int {
	ans := []int{}

	for i := 1; i < int(math.Pow(float64(n), 0.5))+1; i++ {
		if n%i == 0 {
			if n/i != i {
				ans = append(ans, n/i)
			}
			ans = append(ans, i)
		}
	}

	return ans
}

func isThree(n int) bool {
	x := divisors(n)
	return len(x) == 3
}
