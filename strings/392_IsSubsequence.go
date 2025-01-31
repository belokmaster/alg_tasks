package main

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	indexFirst := 0
	for i := range t {
		if t[i] == s[indexFirst] {
			indexFirst++
		}

		if len(s) == indexFirst {
			return true
		}
	}

	return false
}
