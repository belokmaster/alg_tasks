package main

import "sort"

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]
	answer := []rune{}

	for i := range min(len(first), len(last)) {
		if first[i] != last[i] {
			return string(answer)
		}
		answer = append(answer, rune(first[i]))
	}

	return string(answer)
}
