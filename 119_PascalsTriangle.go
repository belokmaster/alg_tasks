package main

func generatePascal(numRows int) [][]int {
	var answer [][]int
	answer = append(answer, []int{1})
	if numRows == 1 {
		return answer
	}

	answer = append(answer, []int{1, 1})
	if numRows == 2 {
		return answer
	}

	for i := 2; i < numRows; i++ {
		dig := make([]int, i+1)
		dig[0] = 1
		dig[i] = 1
		for j := 1; j < len(dig)-1; j++ {
			dig[j] = answer[i-1][j] + answer[i-1][j-1]
		}
		answer = append(answer, dig)
	}
	return answer
}

func getRow(rowIndex int) []int {
	return generate(rowIndex + 1)[rowIndex]
}
