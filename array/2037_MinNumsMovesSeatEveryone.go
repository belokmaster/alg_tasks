package main

import (
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	ans := 0
	for i := 0; i < len(seats); i++ {
		ans += int(math.Abs(float64(students[i] - seats[i])))
	}
	return ans
}
