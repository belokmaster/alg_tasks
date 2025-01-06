package main

// Функция для генерации всех подмножеств
func subsets(nums []int) [][]int {
	// Инициализируем результат с пустым подмножеством
	// `result` изначально содержит одно подмножество — пустое: `{{}}`
	result := [][]int{{}}

	// Проходим по каждому числу в исходном массиве nums
	for _, num := range nums {
		// Создаём временный список для новых подмножеств, которые включают num
		newSubsets := [][]int{}

		for _, subset := range result {
			// Создаём новое подмножество, копируя текущее
			newSubset := append([]int{}, subset...)
			// Добавляем текущий элемент num в новое подмножество
			newSubset = append(newSubset, num)
			// Добавляем это новое подмножество в список newSubsets
			newSubsets = append(newSubsets, newSubset)
		}

		// Добавляем все новые подмножества из newSubsets в общий результат
		result = append(result, newSubsets...)
	}
	return result
}

/*
func main() {
	nums := []int{1, 2, 3}
	subsetsResult := subsets(nums)
	fmt.Println(subsetsResult)
}
*/
