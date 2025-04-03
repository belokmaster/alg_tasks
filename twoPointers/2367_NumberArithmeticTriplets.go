package main

func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[j]-nums[i] == diff && nums[k]-nums[j] == diff {
					ans++
				}
			}
		}
	}

	return ans
}

func arithmeticTripletsGood(nums []int, diff int) int {
	seen := make(map[int]bool)
	ans := 0

	for _, num := range nums {
		if seen[num-diff] && seen[num-2*diff] {
			ans++
		}
		seen[num] = true
	}

	return ans
}
