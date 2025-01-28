package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		x := target - num
		if id, ok := hashMap[x]; ok {
			return []int{id, i}
		}
		hashMap[num] = i
	}

	return []int{0, 1}
}

// через массивы
func findDigit1(nums []int, target, index int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && i != index {
			return i
		}
	}
	return -1
}

func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		pairIndex := findDigit1(nums, target-nums[i], i)
		if pairIndex != -1 {
			return []int{i, pairIndex}
		}
	}
	return []int{0, 1}
}
