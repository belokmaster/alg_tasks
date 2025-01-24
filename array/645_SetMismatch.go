package main

func findErrorNums(nums []int) []int {
	n := len(nums)
	seen := make([]bool, n+1)
	answer := []int{0, 0}

	// цикл нахождения дубля
	for _, num := range nums {
		if seen[num] {
			answer[0] = num
		} else {
			seen[num] = true
		}
	}

	// цикл для пропуска
	for i := 1; i <= n; i++ {
		if !seen[i] {
			answer[1] = i
			break
		}
	}

	return answer
}
