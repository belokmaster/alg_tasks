package main

func balancedStringSplit(s string) int {
	countR := 0
	countL := 0
	ans := 0

	for _, w := range s {
		word := string(w)
		if word == "R" {
			countR++
		} else {
			countL++
		}

		if countL == countR {
			ans++
			countL = 0
			countR = 0
		}
	}

	return ans
}
