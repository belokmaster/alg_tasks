package main

func consistWord(s string, x byte) bool {
	for _, w := range s {
		if byte(w) == x {
			return true
		}
	}
	return false
}

func findWordsContaining(words []string, x byte) []int {
	ans := []int{}
	for i, w := range words {
		if consistWord(w, x) {
			ans = append(ans, i)
		}
	}
	return ans
}

func findWordsContaining1(words []string, x byte) []int {
	var ans []int
	for i, word := range words {
		for _, c := range word {
			if c == rune(x) {
				ans = append(ans, i)
				break
			}
		}
	}

	return ans
}
