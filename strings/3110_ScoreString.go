package main

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -1 * x
	}
}

func scoreOfString(s string) int {
	count := 0
	for i := 0; i < len(s)-1; i++ {
		x, y := int(s[i]), int(s[i+1])
		count += abs(y - x)
	}
	return count
}
