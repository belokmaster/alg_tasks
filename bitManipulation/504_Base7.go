package main

func reverseString2(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes)
}

func convertation(num int) string {
	digits := map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
	}

	var ans string
	for num > 0 {
		d := num % 7
		ans += digits[d]
		num /= 7
	}

	return reverseString2(ans)
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	if num > 0 {
		return convertation(num)
	} else {
		num *= -1
		ans := convertation(num)
		return "-" + ans
	}
}
