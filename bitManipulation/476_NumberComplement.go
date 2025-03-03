package main

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	ans := 0
	mask := 1

	for num > 0 {
		if num&1 == 0 {
			ans += mask
		}
		mask <<= 1
		num >>= 1
	}

	return ans
}
