package main

import "strings"

func countValidWords(sentence string) int {
	res, tokens := 0, strings.Fields(sentence)

	for _, token := range tokens {
		if valid(token) {
			res++
		}
	}

	return res
}

func valid(t string) bool {
	if t[0] == '-' || t[len(t)-1] == '-' {
		return false
	}

	punc := map[rune]bool{'!': true, '.': true, ',': true}
	hCount := 0

	for k, v := range t {
		if v == '-' {
			hCount++

			next := rune(t[k+1])

			if hCount > 1 || punc[next] {
				return false
			}
		}

		if v >= '0' && v <= '9' {
			return false
		}

		if punc[v] && k < len(t)-1 {
			return false
		}
	}

	return true
}
