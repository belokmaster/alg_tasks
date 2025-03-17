package main

func checkIfPangram(sentence string) bool {
	words := make(map[byte]int)
	for _, w := range sentence {
		words[byte(w)]++
	}

	byteWord := byte('a')
	for i := 0; i < 26; i++ {
		_, exists := words[byteWord]
		if exists {
			byteWord++
		} else {
			return false
		}
	}

	return true
}
