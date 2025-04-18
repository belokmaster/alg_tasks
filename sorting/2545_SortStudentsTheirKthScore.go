package main

import "sort"

func sortTheStudents(score [][]int, k int) [][]int {
	type scores struct {
		ind  int
		mark int
	}

	n := len(score)
	m := len(score[0])
	ind := make([]scores, n)
	ans := [][]int{}

	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, m))
		ind[i] = scores{i, score[i][k]}
	}

	sort.Slice(ind, func(i, j int) bool {
		return ind[i].mark > ind[j].mark
	})

	for i, w := range ind {
		ans[i] = score[w.ind]
	}

	return ans
}
