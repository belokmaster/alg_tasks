package main

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes)
}

func convertPositive(num int) string {
	if num == 0 {
		return "0"
	}

	HexDigit := map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "a",
		11: "b",
		12: "c",
		13: "d",
		14: "e",
		15: "f",
	}

	var ans string
	for num > 0 {
		d := num % 16
		ans += HexDigit[d]
		num /= 16
	}

	return reverseString(ans)
}

func reverseBits1(s string) string {
	for len(s) < 8 {
		s += "0"
	}

	hexDigit := map[string]string{
		"0": "f",
		"1": "e",
		"2": "d",
		"3": "c",
		"4": "b",
		"5": "a",
		"6": "9",
		"7": "8",
		"8": "7",
		"9": "6",
		"a": "5",
		"b": "4",
		"c": "3",
		"d": "2",
		"e": "1",
		"f": "0",
	}

	var ans string
	for i := 0; i < 8; i++ {
		ans += hexDigit[string(s[i])]
	}

	return reverseString(ans)
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	var ans string
	if num > 0 {
		ans = convertPositive(num)
		return ans
	}

	num += 1 // прибавляем единицу для того чтобы не сломать инвентирование. в обратном коде нету -0
	// далее инвентируем все биты
	num *= -1
	ans = convertPositive(num)

	ans = reverseString(ans)

	ans = reverseBits1(ans)
	return ans
}
