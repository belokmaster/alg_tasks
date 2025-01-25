package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	// Проходим треугольник с предпоследнего ряда к первому
	for i := len(triangle) - 2; i >= 0; i-- {
		// Обновляем элементы текущего ряда, добавляя минимальное из соседей строки ниже
		for j := 0; j < len(triangle[i]); j++ {
			// Добавляем к текущему элементу минимальное значение из двух соседей снизу
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}
