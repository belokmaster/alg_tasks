package main

import "strconv"

func sumDigits(s string) string {
	var mp = map[rune]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5,
		'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
		'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15,
		'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20,
		'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
	}

	res := ""
	for _, w := range s {
		res += strconv.Itoa(mp[w])
	}
	return res
}

func getLucky(s string, k int) int {
	s = sumDigits(s)
	for i := 0; i < k; i++ {
		sum := 0
		for _, ch := range s {
			sum += int(ch - '0')
		}
		s = strconv.Itoa(sum)
	}
	ans, _ := strconv.Atoi(s)
	return ans
}
