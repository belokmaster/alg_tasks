package main

func countQuadruplets(nums []int) int {
	ans := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				for d := k + 1; d < n; d++ {
					if i < j && i < k && i < d && j < k && j < d && k < d && (nums[i]+nums[j]+nums[k] == nums[d]) {
						ans++
					}
				}
			}
		}
	}

	return ans
}
