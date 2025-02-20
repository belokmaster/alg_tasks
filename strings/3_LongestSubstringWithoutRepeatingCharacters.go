package main

func count(s string, ch byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == ch {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	bestAns := ""
	curAns := ""

	for i := 0; i < len(s); i++ {
		if count(curAns, s[i]) {
			if len(curAns) > len(bestAns) {
				bestAns = curAns
			}

			pos := 0
			for j := 0; j < len(curAns); j++ {
				if curAns[j] == s[i] {
					pos = j + 1
					break
				}
			}
			curAns = curAns[pos:]
		}

		curAns += string(s[i])

		if len(curAns) > len(bestAns) {
			bestAns = curAns
		}
	}

	return len(bestAns)
}
