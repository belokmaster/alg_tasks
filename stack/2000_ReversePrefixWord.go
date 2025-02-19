package main

func reversePrefix(word string, ch byte) string {
	stack := []string{}
	flag := true
	ind := 0

	for i, w := range word {
		if byte(w) == ch {
			flag = false
			stack = append(stack, string(w))
			ind = i
			break
		} else {
			stack = append(stack, string(w))
		}
	}

	lastWord := ""
	for i := len(stack) - 1; i >= 0; i-- {
		lastWord += string(stack[i])
	}

	if flag == true {
		return word
	} else {
		return lastWord + word[ind+1:]
	}
}
