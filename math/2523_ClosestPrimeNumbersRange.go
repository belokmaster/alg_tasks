package main

import "math"

func isPrimeDigit(num int) bool {
	if num == 2 {
		return true
	}

	if num == 1 || num%2 == 0 {
		return false
	}

	for i := 3; i < int(math.Pow(float64(num), 0.5))+1; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func closestPrimes(left int, right int) []int {
	primesDigits := []int{}
	for i := left; i <= right; i++ {
		if isPrimeDigit(i) {
			primesDigits = append(primesDigits, i)
		}
	}

	ans := []int{-1, -1}
	ansDiff := 999999
	for i := 1; i < len(primesDigits); i++ {
		diff := primesDigits[i] - primesDigits[i-1]
		if diff < ansDiff {
			ans[0] = primesDigits[i-1]
			ans[1] = primesDigits[i]
			ansDiff = diff
		}
	}

	return ans
}
