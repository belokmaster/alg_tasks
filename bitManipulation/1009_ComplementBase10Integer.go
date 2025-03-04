package main

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	ans := 0
	mask := 1

	for n > 0 {
		if n&1 == 0 {
			ans += mask
		}
		mask <<= 1
		n >>= 1
	}

	return ans
}
