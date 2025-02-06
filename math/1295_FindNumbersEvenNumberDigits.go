package main

func countDigits(n int) int {
	ans := 0

	for n > 0 {
		ans++
		n /= 10
	}

	return ans
}

func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		sum := countDigits(num)
		if sum%2 == 0 {
			ans++
		}
	}
	return ans
}
