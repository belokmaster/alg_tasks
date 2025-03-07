package main

import "sort"

func frequencySort(s string) string {
	hashMap := make(map[rune]int)
	for _, char := range s {
		hashMap[char]++
	}

	chars := make([]rune, 0, len(hashMap))
	for k := range hashMap {
		chars = append(chars, k)
	}

	sort.Slice(chars, func(i, j int) bool {
		return hashMap[chars[i]] > hashMap[chars[j]]
	})

	ans := make([]rune, 0, len(s))
	for _, char := range chars {
		count := hashMap[char]
		for i := 0; i < count; i++ {
			ans = append(ans, char)
		}
	}

	return string(ans)
}
