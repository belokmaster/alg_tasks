package main

func countConsistentStrings(allowed string, words []string) int {
	wordAllowed := make(map[byte]bool)
	for _, w := range allowed {
		wordAllowed[byte(w)] = true
	}
	ans := 0

	for _, w := range words {
		fl := true

		wordsMap := make(map[byte]bool)
		for _, word := range w {
			wordsMap[byte(word)] = true
		}

		for key := range wordsMap {
			if !wordAllowed[key] {
				fl = false
			}
		}

		if fl {
			ans++
		}
	}

	return ans
}
