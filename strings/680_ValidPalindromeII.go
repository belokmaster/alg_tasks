package main

func isPalindrome1(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func skipIndex(s string, index int) string {
	return s[:index] + s[index+1:]
}

func validPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			if isPalindrome1(skipIndex(s, i)) {
				return true
			}

			if isPalindrome1(skipIndex(s, n-1-i)) {
				return true
			}

			return false
		}
	}
	return true
}
