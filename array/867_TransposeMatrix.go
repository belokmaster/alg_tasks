package main

func transpose(matrix [][]int) [][]int {
	var ans = make([][]int, len(matrix[0]))

	for key := range ans {
		ans[key] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ans[j][i] = matrix[i][j]
		}
	}

	return ans

}
