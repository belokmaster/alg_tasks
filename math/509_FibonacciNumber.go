package main

func fib(n int) int {
	prev, next := 0, 1
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 0; i < n; i++ {
			prev, next = next, next+prev
		}
	}
	return prev
}
