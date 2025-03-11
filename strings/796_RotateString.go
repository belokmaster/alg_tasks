package main

func rotateString(s string, goal string) bool {
	n := len(s)
	newString := s

	for i := 0; i < n; i++ {
		newString = newString[1:] + string(newString[0])

		if newString == goal {
			return true
		}
	}

	return false
}
