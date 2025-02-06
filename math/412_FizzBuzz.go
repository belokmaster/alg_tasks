package main

import "strconv"

func fizzBuzz(n int) []string {
	ans := make([]string, n)

	for i := 0; i < n; i++ {
		x := i + 1
		if x%3 == 0 {
			if x%5 == 0 {
				ans[i] = "FizzBuzz"
			} else {
				ans[i] = "Fizz"
			}
		} else if x%5 == 0 {
			ans[i] = "Buzz"
		} else {
			ans[i] = strconv.Itoa(x)
		}
	}

	return ans
}
