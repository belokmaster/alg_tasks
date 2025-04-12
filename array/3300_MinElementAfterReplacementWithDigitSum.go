package main

func minElement1(nums []int) int {
	ans := int(^uint(0) >> 1)
	for _, num := range nums {
		s := 0
		for num > 0 {
			s += num % 10
			num /= 10
		}
		if s == 1 {
			return 1
		}
		if s < ans {
			ans = s
		}
	}
	return ans
}
