package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap1 := make(map[rune]rune)
	hashMap2 := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		r1 := rune(s[i])
		r2 := rune(t[i])

		val1, ok := hashMap1[r1]
		if !ok {
			hashMap1[r1] = r2
			val1 = r2
		}

		if val1 != r2 {
			return false
		}

		val2, ok := hashMap2[r2]
		if !ok {
			hashMap2[r2] = r1
			val2 = r1
		}

		if val2 != r1 {
			return false
		}
	}

	return true
}

// брут форс жесточайший
func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[string]string)
	for i, w := range t {
		hashMap[string(s[i])] = string(w)
	}

	values := []string{}
	for _, value := range hashMap {
		for i := 0; i < len(values); i++ {
			if value == values[i] {
				return false
			}
		}
		values = append(values, value)
	}

	ans := ""
	for i := 0; i < len(s); i++ {
		ans += (hashMap[string(s[i])])
	}

	return ans == t
}
