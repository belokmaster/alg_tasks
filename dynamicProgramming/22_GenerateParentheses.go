package main

func generateParenthesis(n int) []string {
	ans := make([][]string, n+1)
	ans[0] = []string{""}

	for k := 1; k <= n; k++ {
		ans[k] = make([]string, 0)
		for i := 0; i < k; i++ {
			for _, left := range ans[i] {
				for _, right := range ans[k-i-1] {
					ans[k] = append(ans[k], "("+left+")"+right)
				}
			}
		}
	}

	return ans[n]
}
