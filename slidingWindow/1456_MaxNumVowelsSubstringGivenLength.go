package main

func maxVowels(s string, k int) int {
	ans := 0
	count := 0

	vowels := make(map[byte]bool)
	vowels['a'] = true
	vowels['e'] = true
	vowels['i'] = true
	vowels['o'] = true
	vowels['u'] = true

	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			count++
		}
	}

	ans = count
	if ans == k {
		return ans
	}

	for i := k; i < len(s); i++ {
		if vowels[s[i]] {
			count++
		}

		// s[i-k] — это самый левый символ, который выходит из окна
		if vowels[s[i-k]] {
			count--
		}

		if count > ans {
			ans = count
		}
	}

	return ans
}
