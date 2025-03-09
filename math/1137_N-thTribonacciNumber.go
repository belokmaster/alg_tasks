package main

func tribonacci(n int) int {
	x, y, z := 0, 1, 1
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		for i := 0; i < n; i++ {
			x, y, z = y, z, x+y+z
		}
	}
	return x
}
