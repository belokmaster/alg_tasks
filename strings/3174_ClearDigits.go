package main

import "unicode"

func digitInd(s string) int {
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			return i
		}
	}
	return -1
}

// конкатенация строк дорогая операция. лучше использовать стек
func clearDigits(s string) string {
	x := digitInd(s)
	for x != -1 {
		s = s[:x-1] + s[x+1:]
		x = digitInd(s)
	}
	return s
}
