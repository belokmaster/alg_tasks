package main

func toLowerCase(s string) string {
	ans := ""

	for _, w := range s {
		if w >= 65 && w <= 90 {
			ans += string(w + 32)
		} else {
			ans += string(w)
		}
	}

	return ans
}
