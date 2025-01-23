package main

// search ищет заданное значение target в отсортированном и возможно повёрнутом массиве nums.
// Возвращает индекс найденного элемента или -1, если элемент не найден.
func search(nums []int, target int) int {
	// Устанавливаем начальные границы поиска
	left := 0
	right := len(nums) - 1

	// Пока левая граница меньше или равна правой, продолжаем поиск
	for left <= right {
		// Находим середину текущего диапазона
		mid := (left + right) / 2

		// Если элемент в середине равен target, возвращаем его индекс
		if nums[mid] == target {
			return mid
		}

		// Проверяем, отсортирована ли левая часть массива
		if nums[left] <= nums[mid] {
			// Если левая часть отсортирована, проверяем, лежит ли target в её пределах
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Сужаем поиск до левой части
			} else {
				left = mid + 1 // Иначе ищем в правой части
			}
		} else {
			// Если левая часть не отсортирована, значит отсортирована правая
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Сужаем поиск до правой части
			} else {
				right = mid - 1 // Иначе ищем в левой части
			}
		}
	}

	return -1
}

/*
func main() {
	// Примеры использования функции search
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	result := search(nums, target)
	fmt.Printf("Индекс элемента %d: %d\n", target, result)

	// Пример с отсутствующим элементом
	target = 3
	result = search(nums, target)
	fmt.Printf("Индекс элемента %d: %d\n", target, result)

	// Пример с неповернутым массивом
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	target = 5
	result = search(nums, target)
	fmt.Printf("Индекс элемента %d: %d\n", target, result)
}
*/
