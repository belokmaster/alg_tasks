package main

// через мапу, где отмечается, была ли цифра или нет
func findDisappearedNumbers(nums []int) []int {
	hash := make([]bool, len(nums))

	for _, num := range nums {
		hash[num-1] = true
	}

	ans := []int{}
	for i := range hash {
		if !hash[i] {
			ans = append(ans, i+1)
		}
	}

	return ans
}

// брут форс через мапу
func findDisappearedNumbers1(nums []int) []int {
	hashMap := make(map[int]int)

	arr := []int{}
	for i := 0; i < len(nums); i++ {
		arr = append(arr, i+1)
	}

	for _, num := range arr {
		hashMap[num]++
	}

	for _, num := range nums {
		hashMap[num]++
	}

	ans := []int{}

	for key, value := range hashMap {
		if value == 1 {
			ans = append(ans, key)
		}
	}

	return ans
}
