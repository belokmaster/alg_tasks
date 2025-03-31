package main

func valueAfterKSeconds(n int, k int) int {
	mod := 1000000007

	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}

	for i := 1; i <= k; i++ {
		for j := 1; j < n; j++ {
			arr[j] = (arr[j-1] + arr[j]) % mod
		}
	}

	return arr[n-1]
}
