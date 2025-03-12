package main

func numberOfSubstrings(s string) int {
	count := [3]int{}
	left := 0
	ans := 0

	for right := 0; right < len(s); right++ {
		count[s[right]-'a']++
		for count[0] > 0 && count[1] > 0 && count[2] > 0 {
			ans += len(s) - right
			count[s[left]-'a']--
			left++
		}
	}

	return ans
}
