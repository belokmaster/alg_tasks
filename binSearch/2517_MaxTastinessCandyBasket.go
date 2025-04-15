package main

import "sort"

func canPick(price []int, k int, minDiff int) bool {
	count := 1
	last := price[0]

	for i := 1; i < len(price); i++ {
		if price[i]-last >= minDiff {
			count++
			last = price[i]
			if count == k {
				return true
			}
		}
	}

	return false
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	ans := 0
	left := 0
	right := price[len(price)-1] - price[0]

	for left <= right {
		mid := (right + left) / 2
		if canPick(price, k, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
