package main

import "sort"

// решение через мапу, ее сортировку ключей по значением и вывод.
// проще через хипу делать конечно
func topKFrequent(nums []int, k int) []int {
	countDigits := make(map[int]int)
	sort.Ints(nums)
	for _, num := range nums {
		countDigits[num]++
	}

	keys := make([]int, 0, len(countDigits))
	for key := range countDigits {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return countDigits[keys[i]] > countDigits[keys[j]]
	})

	ans := []int{}
	for i := 0; i < k; i++ {
		ans = append(ans, keys[i])
	}

	return ans
}
