package main

import "unicode"

func toString(s string) string {
	ans := []rune{}

	for _, w := range s {
		if unicode.IsLetter(w) || unicode.IsDigit(w) {
			ans = append(ans, unicode.ToLower(w))
		}
	}

	return string(ans)
}

func isPalindrome(s string) bool {
	words := toString(s)

	for i := 0; i < len(words)/2; i++ {
		if words[i] != words[len(words)-1-i] {
			return false
		}
	}

	return true
}
