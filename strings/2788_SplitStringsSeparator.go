package main

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := []string{}

	for _, w := range words {
		word := ""
		for i := 0; i < len(w); i++ {
			if w[i] != separator {
				word += string(w[i])
			} else {
				if word != "" {
					ans = append(ans, word)
					word = ""
				}
			}
		}
		if word != "" {
			ans = append(ans, word)
		}
	}

	return ans
}
