package main

func countWord(s string) int {
	count := 0
	for _, w := range s {
		if string(w) == " " {
			count++
		}
	}
	return count + 1
}

func mostWordsFound(sentences []string) int {
	ans := 0
	for i := 0; i < len(sentences); i++ {
		x := countWord(sentences[i])
		if x > ans {
			ans = x
		}
	}
	return ans
}
