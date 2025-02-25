package main

func detectAllCapital(word string) bool {
	for _, w := range word {
		if !(w >= 65 && w <= 90) {
			return false
		}
	}
	return true
}

func detectAllNotCapital(word string) bool {
	for _, w := range word {
		if w >= 65 && w <= 90 {
			return false
		}
	}
	return true
}
func detectCapitalUse(word string) bool {
	if detectAllCapital(word) {
		return true
	}

	if detectAllNotCapital(word) {
		return true
	}

	if detectAllNotCapital(word[1:]) {
		return true
	}

	return false
}
