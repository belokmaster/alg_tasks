package main

import "strings"

// норм вариант
func removeOccurrences(s string, part string) string {
	index := strings.Index(s, part)
	for index != -1 {
		s = s[:index] + s[index+len(part):]
		index = strings.Index(s, part)
	}
	return s
}

func findOccurance1(s string, part string) int {
	if part == "" {
		return 0
	}

	for i := 0; i <= len(s)-len(part); i++ {
		if s[i:i+len(part)] == part {
			return i
		}
	}

	return -1
}

func removeOccurrences1(s string, part string) string {
	x := findOccurance1(s, part)
	if x == -1 {
		return s
	}

	for x != -1 {
		d := len(part)
		s = s[0:x] + s[x+d:]
		x = findOccurance1(s, part)
	}

	return s
}
