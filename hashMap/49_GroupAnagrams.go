package main

func createRunesArr(s string) []rune {
	runes := []rune{}
	for _, w := range s {
		runes = append(runes, w)
	}
	return runes
}

func sortRunes1(runes []rune) []rune {
	for i := 1; i < len(runes); i++ {
		key := runes[i]
		j := i - 1
		for j >= 0 && runes[j] > key {
			runes[j+1] = runes[j]
			j--
		}
		runes[j+1] = key
	}
	return runes
}

func makeString(runes []rune) string {
	ans := ""
	for _, word := range runes {
		ans += string(word)
	}
	return ans
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string)
	for _, word := range strs {
		runes := createRunesArr(word)
		s := makeString(sortRunes1(runes))
		hashMap[s] = append(hashMap[s], word)
	}

	ans := [][]string{}
	for _, value := range hashMap {
		ans = append(ans, value)
	}

	return ans
}
