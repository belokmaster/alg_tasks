package main

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	var diffs []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffs = append(diffs, i)
		}
	}

	if len(diffs) != 2 {
		return false
	}

	i, j := diffs[0], diffs[1]
	return s1[i] == s2[j] && s1[j] == s2[i]
}
