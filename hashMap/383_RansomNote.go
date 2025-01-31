package main

func canConstruct(ransomNote string, magazine string) bool {
	words := make(map[rune]int)
	for _, w := range magazine {
		words[w]++
	}

	for _, w := range ransomNote {
		_, ok := words[w]
		if !ok || words[w] == 0 {
			return false
		} else {
			words[w]--
		}
	}
	return true
}
