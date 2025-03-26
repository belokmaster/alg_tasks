package main

import (
	"math"
	"strings"
)

func findPermutationDifference(s string, t string) int {
	ans := 0
	for i, w := range s {
		ind := strings.Index(t, string(w))
		x := int(math.Abs(float64(i) - float64(ind)))
		ans += x
	}
	return ans
}
