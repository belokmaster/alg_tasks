package main

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	diagonalMap := make(map[int][]int)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalMap[i+j] = append(diagonalMap[i+j], mat[i][j])
		}
	}

	for k := 0; k < m+n-1; k++ {
		if k%2 == 0 {
			for i := len(diagonalMap[k]) - 1; i >= 0; i-- {
				ans = append(ans, diagonalMap[k][i])
			}
		} else {
			ans = append(ans, diagonalMap[k]...)
		}
	}

	return ans
}
