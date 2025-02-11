package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	return num
}

func guessNumber(n int) int {
	var left int = 0
	var right int = n

	for left <= right {
		mid := (left + right) / 2
		result := guess(mid)
		if result == 0 {
			return mid
		} else if result > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
