package main

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}
