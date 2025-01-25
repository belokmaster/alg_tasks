package main

// mySqrt реализует бинарный поиск для нахождения целой части квадратного корня числа x.
func mySqrt(x int) int {
	left := 0
	right := x + 1
	mid := (left + right) / 2

	for left <= right {
		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}

		mid = (left + right) / 2
	}

	return mid
}
