package main

func numberOfSteps(num int) int {
	ans := 0

	for num > 0 {
		if num&1 == 0 {
			num /= 2
			ans++
		} else {
			num -= 1
			ans++
		}
	}

	return ans
}
