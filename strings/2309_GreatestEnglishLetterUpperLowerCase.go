package main

import "sort"

func greatestLetter(s string) string {
	wordsMap := make(map[rune]bool)
	ans := []string{}

	for _, w := range s {
		wordsMap[rune(w)] = true
	}

	for key := range wordsMap {
		if key >= 65 && key <= 90 {
			if wordsMap[key+32] {
				ans = append(ans, string(key))
			}
		}
	}

	sort.Strings(ans)
	if len(ans) == 0 {
		return ""
	}
	return ans[len(ans)-1]
}
