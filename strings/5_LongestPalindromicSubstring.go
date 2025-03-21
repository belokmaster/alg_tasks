package main

func isPalindrom(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func longestPalindrome(s string) string {
	n := len(s)
	ans := string(s[0])
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			x := string(s[i:j])
			if isPalindrom(x) {
				if len(x) > len(ans) {
					ans = x
				}
			}
		}
	}
	return ans
}
