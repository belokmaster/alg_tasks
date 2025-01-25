package main

/*
	Для двух строк s и t мы говорим "t делит s" тогда и только тогда,
	когда s = t + t + t + ... + t + t (т.е. t соединяется само с собой один или несколько раз).
	Учитывая две строки str1 и str2, верните наибольшую строку x таким образом,
	чтобы x делило как str1, так и str2.
*/

// строка s - конкатенация одной или нескольких копий строки x.\
// ищем нод для поиска длины искомой x
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	gcdStrings := gcd(len(str1), len(str2))

	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:gcdStrings]
}
