package main

import (
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	ans := 0
	isNegative := false
	const intMax = int64(1<<31 - 1)
	const intMin = int64(-1 << 31)

	for i, sym := range s {
		if i == 0 && sym == '-' {
			isNegative = true
			continue
		}

		if i == 0 && sym == '+' {
			continue
		}

		if ans == 0 && string(sym) == "0" {
			continue
		}

		if !unicode.IsDigit(sym) {
			break
		}

		ans = ans*10 + int(sym-'0')

		if !isNegative && int64(ans) > intMax {
			return int(intMax)
		}
		if isNegative && int64(-ans) < intMin {
			return int(intMin)
		}
	}

	if isNegative {
		ans = -ans
	}

	return ans
}
