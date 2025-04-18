package main

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay(n - 1)
	ans := ""
	count := 1

	for i := 1; i < len(prev); i++ {
		if prev[i] == prev[i-1] {
			count++
		} else {
			ans += strconv.Itoa(count) + string(prev[i-1])
			count = 1
		}
	}
	ans += strconv.Itoa(count) + string(prev[len(prev)-1])

	return ans
}
