package main

func sortRunes(arr []rune) []rune {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	words1 := []rune{}
	words2 := []rune{}

	for _, word := range s {
		words1 = append(words1, word)
	}

	for _, word := range t {
		words2 = append(words2, word)
	}

	sortRunes(words1)
	sortRunes(words2)

	for i := 0; i < len(s); i++ {
		if words1[i] != words2[i] {
			return false
		}
	}

	return true
}
