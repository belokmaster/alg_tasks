package main

// shiftingLetters применяет смещения к строке s на основе заданных операций shifts.
// s — исходная строка, shifts — массив операций, где каждая операция задаётся
// диапазоном [start, end] и направлением (1 — вправо, 0 — влево).
func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n+1) // Массив для накопления изменений

	// Обрабатываем каждую операцию и обновляем массив diff
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			diff[start]++ // Увеличиваем смещение в начале диапазона
			diff[end+1]-- // Уменьшаем смещение за пределами диапазона
		} else {
			diff[start]-- // Уменьшаем смещение в начале диапазона
			diff[end+1]++ // Увеличиваем смещение за пределами диапазона
		}
	}

	currentShift := 0         // Накопленное смещение
	result := make([]byte, n) // Результирующая строка

	// Применяем накопленные смещения к каждому символу
	for i := 0; i < n; i++ {
		currentShift += diff[i]                        // Обновляем текущее смещение
		shifted := (int(s[i]-'a') + currentShift) % 26 // Сдвигаем символ
		if shifted < 0 {
			shifted += 26 // Обработка отрицательных смещений
		}
		result[i] = byte(shifted) + 'a' // Преобразуем обратно в символ
	}

	return string(result)
}

/*
func main() {
	// Пример 1: простая строка и смещения
	s := "abc"
	shifts := [][]int{
		{0, 1, 1},
		{1, 2, 0},
	}
	fmt.Println("Результат:", shiftingLetters(s, shifts)) // Ожидаемый результат: "ace"

	// Пример 2: строка с большими диапазонами
	s = "xyz"
	shifts = [][]int{
		{0, 2, 1},
	}
	fmt.Println("Результат:", shiftingLetters(s, shifts)) // Ожидаемый результат: "yza"

	// Пример 3: без операций
	s = "hello"
	shifts = [][]int{}
	fmt.Println("Результат:", shiftingLetters(s, shifts)) // Ожидаемый результат: "hello"
}
*/
