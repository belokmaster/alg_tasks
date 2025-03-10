package main

import (
	"slices"
	"sort"
	"strconv"
)

// ДИЧАЙШИЙ ГОВНОКОД НО ЗАТО РАБОТАЕТ
func findRelativeRanks(score []int) []string {
	n := len(score)
	ans := make([]string, n)

	hashMap := make(map[int]int)
	for i := range n {
		hashMap[score[i]] = i
	}

	sortedScore := make([]int, n)
	copy(sortedScore, score)
	sort.Ints(sortedScore)
	slices.Reverse(sortedScore)

	place := 4

	for i := range n {
		if i == 0 {
			ans[hashMap[sortedScore[i]]] = "Gold Medal"
		} else if i == 1 {
			ans[hashMap[sortedScore[i]]] = "Silver Medal"
		} else if i == 2 {
			ans[hashMap[sortedScore[i]]] = "Bronze Medal"
		} else {
			stringPlace := strconv.Itoa(place)
			place++
			ans[hashMap[sortedScore[i]]] = stringPlace
		}
	}
	return ans
}
