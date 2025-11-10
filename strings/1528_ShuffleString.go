package main

func restoreString(s string, indices []int) string {
	hashMap := make(map[int]rune)
	for i, w := range s {
		hashMap[indices[i]] = w
	}

	ans := []rune{}
	for i := 0; i < len(indices); i++ {
		ans = append(ans, hashMap[i])
	}

	return string(ans)
}
