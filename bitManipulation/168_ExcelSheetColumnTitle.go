package main

func reverseString4(s string) string {
	n := len(s)
	runes := []rune(s)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func convertToTitle(columnNumber int) string {
	hashMap := map[int]string{
		1: "A", 2: "B", 3: "C", 4: "D", 5: "E",
		6: "F", 7: "G", 8: "H", 9: "I", 10: "J",
		11: "K", 12: "L", 13: "M", 14: "N", 15: "O",
		16: "P", 17: "Q", 18: "R", 19: "S", 20: "T",
		21: "U", 22: "V", 23: "W", 24: "X", 25: "Y",
		26: "Z",
	}

	var x int
	var ans string
	for columnNumber > 0 {
		columnNumber-- // нет 0, поэтому приходится так делать. основной прикол в этом
		x = columnNumber % 26
		ans += hashMap[x+1]
		columnNumber /= 26
	}

	ans = reverseString4(ans)
	return ans
}
