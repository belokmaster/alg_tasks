package main

func hammingWeight(num int) int {
	ans := 0
	for num != 0 {
		if num%2 == 1 {
			ans++
		}
		num /= 2
	}
	return ans
}
