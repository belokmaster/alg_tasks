package main

func arrangeCoins(n int) int {
	left := 0
	right := n

	for left <= right {
		mid := left + (right-left)/2
		curSum := mid * (mid + 1) / 2
		if curSum == n {
			return mid
		} else if curSum < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}
