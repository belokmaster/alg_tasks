package main

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left := matrix[0][0]      // Минимальное возможное значение (левый верхний элемент)
	right := matrix[n-1][n-1] // Максимальное возможное значение (правый нижний элемент)

	for left < right {
		mid := left + (right-left)/2
		count := 0 // Количество элементов, меньших или равных mid

		// Подсчет количества элементов, меньших или равных mid, по строкам
		for i := 0; i < n; i++ {
			j := n - 1 // Начинаем с последнего элемента строки

			// Сдвигаемся влево, пока элемент больше mid
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += j + 1 // Добавляем количество элементов <= mid в текущей строке
		}

		// Если количество таких элементов меньше k, значит ответ лежит справа
		if count < k {
			left = mid + 1
		} else {
			// Иначе ответ лежит в левой половине (включая mid)
			right = mid
		}
	}

	// Возвращаем left или right (они равны на этом этапе)
	return right
}

/*
	func kthSmallest(matrix [][]int, k int) int {
    digits := []int{}
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            digits = append(digits, matrix[i][j])
        }
    }
    sort.Ints(digits)
    return digits[k - 1]
	}
*/
