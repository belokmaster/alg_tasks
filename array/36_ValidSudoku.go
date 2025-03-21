package main

func checkMatrix(arr [][]byte) bool {
	boolMap := make(map[byte]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if arr[i][j] == 46 {
				continue
			}
			if boolMap[arr[i][j]] {
				return true
			}
			boolMap[arr[i][j]] = true
		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	n := 9
	for i := 0; i < n; i++ {
		boolMap := make(map[byte]bool)
		for j := 0; j < n; j++ {
			if board[i][j] == 46 {
				continue
			}
			if boolMap[board[i][j]] {
				return false
			}
			boolMap[board[i][j]] = true
		}
	}

	for i := 0; i < n; i++ {
		boolMap := make(map[byte]bool)
		for j := 0; j < n; j++ {
			if board[j][i] == 46 {
				continue
			}
			if boolMap[board[j][i]] {
				return false
			}
			boolMap[board[j][i]] = true
		}
	}

	matrix1 := [][]byte{}
	matrix2 := [][]byte{}
	matrix3 := [][]byte{}

	for i := 0; i < n; i++ {
		arr := board[i]

		matrix1 = append(matrix1, arr[0:3])
		matrix2 = append(matrix2, arr[3:6])
		matrix3 = append(matrix3, arr[6:9])

		if len(matrix1) == 3 {
			if checkMatrix(matrix1) || checkMatrix(matrix2) || checkMatrix(matrix3) {
				return false
			}
			matrix1 = [][]byte{}
			matrix2 = [][]byte{}
			matrix3 = [][]byte{}
		}

	}

	return true
}
