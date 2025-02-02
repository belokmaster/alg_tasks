package main

import (
	"strconv"
	"strings"
	"unicode"
)

func areNumbersAscending(s string) bool {
	digits := []int{}
	parts := strings.Fields(s)

	for _, part := range parts {
		isNumber := true
		for _, r := range part {
			if !unicode.IsDigit(r) {
				isNumber = false
				break
			}
		}
		if isNumber {
			x, _ := strconv.Atoi(part)
			digits = append(digits, x)
		}
	}

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] >= digits[i+1] {
			return false
		}
	}

	return true
}
