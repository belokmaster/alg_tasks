package main

// бинарный поиск для нахождения квадратного корня
func mySqrt1(x int) int {
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

func isPerfectSquare(num int) bool {
	sqrt := mySqrt1(num)
	return sqrt*sqrt == num
}
