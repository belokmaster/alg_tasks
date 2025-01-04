package main

import "fmt"

/* Учитывая целое число numRows, верните первые числа треугольника Паскаля.
   В треугольнике Паскаля каждое число является суммой двух чисел, расположенных непосредственно над ним, как показано на рисунке
*/

// Функция для генерации треугольника Паскаля
func generate(numRows int) [][]int {
	var answer [][]int
	// Первая строка треугольника
	answer = append(answer, []int{1})

	// Если нужно только одну строку, возвращаем результат сразу
	if numRows == 1 {
		return answer
	}

	// Добавляем вторую строку
	answer = append(answer, []int{1, 1})

	// Если нужно только две строки, возвращаем результат
	if numRows == 2 {
		return answer
	}

	// Генерируем оставшиеся строки
	for i := 2; i < numRows; i++ {
		// Создаём новую строку длины i + 1
		dig := make([]int, i+1)
		dig[0] = 1 // первый элемент равен 1
		dig[i] = 1 // последний элемент равен 1

		// Заполняем промежуточные элементы
		for j := 1; j < len(dig)-1; j++ {
			dig[j] = answer[i-1][j] + answer[i-1][j-1]
		}
		answer = append(answer, dig)
	}
	return answer
}

func main() {
	numRows := 5
	result := generate(numRows)
	fmt.Println(result) // Ожидаемый вывод для numRows = 5: [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
}
