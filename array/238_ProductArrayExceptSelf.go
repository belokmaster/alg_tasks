package main

func prodDigits(nums []int) int {
	ans := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			ans *= nums[i]
		}
	}

	return ans
}

func productExceptSelf(nums []int) []int {
	zeroCount := 0
	zeroIndex := -1
	for i, num := range nums {
		if num == 0 {
			zeroCount++
			zeroIndex = i
		}
	}

	ans := make([]int, len(nums))
	if zeroCount > 1 {
		return ans
	}

	sum := prodDigits(nums)
	if zeroCount == 1 {
		ans[zeroIndex] = sum
		return ans
	}

	for i := 0; i < len(nums); i++ {
		ans[i] = sum / nums[i]
	}

	return ans
}
