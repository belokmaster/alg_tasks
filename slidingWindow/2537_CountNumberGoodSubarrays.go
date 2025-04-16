package main

func countGood(nums []int, k int) int64 {
	hashMap := make(map[int]int)
	countPairs := 0
	left := 0
	var ans int64

	for right := 0; right < len(nums); right++ {
		val := nums[right]

		countPairs += hashMap[val]
		hashMap[val]++

		for countPairs >= k {
			ans += int64(len(nums) - right)

			hashMap[nums[left]]--
			countPairs -= hashMap[nums[left]]
			left++
		}
	}

	return ans
}
