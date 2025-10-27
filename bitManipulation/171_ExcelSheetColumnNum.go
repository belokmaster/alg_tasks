package main

import "math"

// лучшее решение.
func titleToNumber(columnTitle string) int {
	sum := 0
	multy := 1

	for i := len(columnTitle) - 1; i >= 0; i-- {
		index := columnTitle[i] - 65
		sum += int(index+1) * multy
		multy *= 26
	}

	return sum
}

// бич решение мое
func reverseString3(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes)
}

func titleToNumber1(columnTitle string) int {
	hashMap := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5,
		"F": 6, "G": 7, "H": 8, "I": 9, "J": 10,
		"K": 11, "L": 12, "M": 13, "N": 14, "O": 15,
		"P": 16, "Q": 17, "R": 18, "S": 19, "T": 20,
		"U": 21, "V": 22, "W": 23, "X": 24, "Y": 25,
		"Z": 26,
	}

	ans := 0
	n := len(columnTitle)

	columnTitle = reverseString3(columnTitle)
	for i := 0; i < n; i++ {
		cat := int(math.Pow(float64(26), float64(i)))
		ans += hashMap[string(columnTitle[i])] * cat
	}

	return ans
}
