package main

func checkString(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if string(s[i]) == "b" && string(s[i+1]) == "a" {
			return false
		}
	}
	return true
}
