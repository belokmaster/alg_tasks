package main

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ans := []int{0, 1}
	base := 2

	for i := 1; i < n; i++ {
		rlen := len(ans)
		for j := rlen - 1; j >= 0; j-- {
			ans = append(ans, ans[j]+base)
		}
		base *= 2
	}

	return ans
}
