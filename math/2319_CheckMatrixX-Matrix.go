package main

func sumDiagonals(grid [][]int) bool {
	n := len(grid[0])
	for i := 0; i < n; i++ {
		if grid[i][i] == 0 || grid[i][n-i-1] == 0 {
			return false
		}
	}
	return true
}

func diagonalSum1(mat [][]int) int {
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

func checkXMatrix(grid [][]int) bool {
	if !sumDiagonals(grid) {
		return false
	}

	n := len(grid[0])
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
		}
	}

	sumDiag := diagonalSum(grid)
	return sum-sumDiag == 0
}
