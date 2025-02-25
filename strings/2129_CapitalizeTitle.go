package main

import (
	"fmt"
	"strings"
)

func capitalizeWord(word string) string {
	ans := ""
	if len(word) <= 2 {
		for i := 0; i < len(word); i++ {
			if word[i] >= 65 && word[i] <= 90 {
				ans += string(word[i] + 32)
			} else {
				ans += string(word[i])
			}
		}
		return ans
	}

	for i, w := range word {
		if i == 0 {
			if w >= 'a' && w <= 'z' {
				ans += string(w - 32)
			} else {
				ans += string(w)
			}
		} else {
			if w >= 'A' && w <= 'Z' {
				ans += string(w + 32)
			} else {
				ans += string(w)
			}
		}
	}
	return ans
}

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	ans := ""
	fmt.Println(words)
	for i := 0; i < len(words); i++ {
		words[i] = capitalizeWord(words[i])
		ans += words[i]
		if i != len(words)-1 {
			ans += " "
		}
	}

	return ans
}
