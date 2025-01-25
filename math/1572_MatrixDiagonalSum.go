package main

func diagonalSum(mat [][]int) int {
	n := len(mat[0])
	primaryDiagonal := 0
	secondaryDiagonal := 0

	for i := 0; i < n; i++ {
		primaryDiagonal += mat[i][i]
		if i != n-1-i {
			secondaryDiagonal += mat[i][n-i-1]
		}
	}

	return primaryDiagonal + secondaryDiagonal
}
