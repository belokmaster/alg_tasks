package main

func alternatingSum(nums []int) int {
	ans := 0
	for i, num := range nums {
		if i%2 == 0 {
			ans += num
		} else {
			ans -= num
		}
	}

	return ans
}
