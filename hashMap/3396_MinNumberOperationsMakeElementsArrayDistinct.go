package main

func isUnique(nums []int) bool {
	hashMap := make(map[int]bool)
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if hashMap[num] {
			return false
		}
		hashMap[num] = true
	}

	return true
}

func minimumOperations(nums []int) int {
	ans := 0
	if isUnique(nums) {
		return ans
	}

	for i := 0; i < len(nums)%3; i++ {
		nums = append(nums, 0)
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		nums = nums[3:]
		ans++
		if isUnique(nums) {
			return ans
		}
	}

	return ans
}
