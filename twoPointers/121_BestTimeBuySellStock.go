package main

func maxProfit(prices []int) int {
	buyPrice := prices[0]
	ans := 0

	for i := 1; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		}
		if prices[i]-buyPrice > ans {
			ans = prices[i] - buyPrice
		}
	}

	return ans
}
