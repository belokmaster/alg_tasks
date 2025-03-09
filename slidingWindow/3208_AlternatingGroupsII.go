package main

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	count := 0
	left := 0

	for right := 0; right < n+k-1; right++ {
		if right > 0 && colors[right%n] == colors[(right-1)%n] {
			left = right
		}

		if right-left+1 >= k {
			count++
		}
	}

	return count
}
