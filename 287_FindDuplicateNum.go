package main

// мапа
func findDuplicate(nums []int) int {
	seen := make([]bool, len(nums))

	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}

// findDuplicate принимает массив nums, содержащий n + 1 целых чисел, где каждое число в диапазоне [1, n].
// Массив гарантированно содержит одно дублирующееся число, которое нужно найти.
// Решение использует алгоритм с двумя указателями (медленный и быстрый) для обнаружения цикла.
func findDuplicate1(nums []int) int {
	// два указателя
	slow := nums[0]
	fast := nums[0]

	// находим цикл через алгоритм флойда
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// начало цикла
	slow = nums[0]
	for slow != fast { // Пока медленный и быстрый указатели не встретятся:
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
