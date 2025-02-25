package main

func numberOfSpecialChars(word string) int {
	wordsMap := make(map[rune]bool)
	ans := 0

	for _, w := range word {
		wordsMap[rune(w)] = true
	}

	for key := range wordsMap {
		if key >= 65 && key <= 90 {
			if wordsMap[key+32] {
				ans++
			}
		}
	}

	return ans
}
