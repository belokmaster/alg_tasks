package main

func oneCount(n int) int {
	ans := 0

	for n > 0 {
		if n%2 == 1 {
			ans++
		}
		n /= 2
	}

	return ans
}

func countBits(n int) []int {
	ans := []int{}

	for i := 0; i <= n; i++ {
		ans = append(ans, oneCount(i))
	}

	return ans
}
