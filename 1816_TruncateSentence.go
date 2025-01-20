package main

import "strings"

func truncateSentence1(s string, k int) string {
	arr := strings.Split(s, " ")
	return strings.Join(arr[:k], " ")
}

func truncateSentence(s string, k int) string {
	words := []string{}
	s += " "
	si := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			words = append(words, si)
			si = ""
		} else {
			si += string(s[i])
		}
	}

	answer := ""
	for i := 0; i < k; i++ {
		if i == k-1 {
			answer += words[i]
			break
		}
		answer += words[i]
		answer += " "
	}
	return answer
}
