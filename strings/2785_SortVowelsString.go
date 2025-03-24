package main

import "sort"

func sortVowels(s string) string {
	vowelsArr := []rune{}
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	for _, w := range s {
		if vowels[w] {
			vowelsArr = append(vowelsArr, w)
		}
	}

	sort.Slice(vowelsArr, func(i, j int) bool {
		return vowelsArr[i] < vowelsArr[j]
	})

	ans := []rune(s)
	index := 0

	for i, w := range s {
		if vowels[w] {
			ans[i] = vowelsArr[index]
			index++
		}
	}

	return string(ans)
}
