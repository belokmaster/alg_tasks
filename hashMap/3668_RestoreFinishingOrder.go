package main

func recoverOrder(order []int, friends []int) []int {
	hashMap := make(map[int]bool)
	for _, v := range friends {
		hashMap[v] = true
	}

	var ans []int
	for _, v := range order {
		if hashMap[v] {
			ans = append(ans, v)
		}
	}

	return ans
}
