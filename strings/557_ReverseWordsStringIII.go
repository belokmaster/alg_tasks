package main

import (
	"slices"
	"strings"
)

func reverseWord(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func reverseWords1(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = reverseWords1(word)
	}
	return strings.Join(words, " ")
}
