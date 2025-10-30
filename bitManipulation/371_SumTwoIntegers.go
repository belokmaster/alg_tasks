package main

func toBin(x int) [13]byte {
	num := [13]byte{}

	if x < 0 {
		x = -x
		for i := 0; i < 12; i++ {
			num[i] = byte(x % 2)
			x /= 2
		}

		carry := byte(1)
		for i := 0; i < 12; i++ {
			bit := 1 - num[i]
			sum := bit + carry
			num[i] = sum % 2
			carry = sum / 2
		}
		num[12] = 1
	} else {
		for i := 0; i < 12; i++ {
			num[i] = byte(x % 2)
			x /= 2
		}
		num[12] = 0
	}

	return num
}

func toInt(num [13]byte) int {
	if num[12] == 1 {
		carry := byte(1)
		temp := [13]byte{}
		for i := 0; i < 12; i++ {
			bit := 1 - num[i]
			sum := bit + carry
			temp[i] = sum % 2
			carry = sum / 2
		}

		x := 0
		for i := 0; i < 12; i++ {
			if temp[i] == 1 {
				x += 1 << i
			}
		}
		return -x
	} else {
		x := 0
		for i := 0; i < 12; i++ {
			if num[i] == 1 {
				x += 1 << i
			}
		}
		return x
	}
}

func addBinary(n1, n2 [13]byte) [13]byte {
	ans := [13]byte{}
	var carry byte

	for i := 0; i < 13; i++ {
		sum := n1[i] + n2[i] + carry
		ans[i] = sum % 2
		carry = sum / 2
	}

	return ans
}

func getSum(a int, b int) int {
	n1 := toBin(a)
	n2 := toBin(b)
	ansByte := addBinary(n1, n2)
	ans := toInt(ansByte)
	return ans
}
