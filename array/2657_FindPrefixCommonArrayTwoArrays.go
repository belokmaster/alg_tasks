package main

func countCommonElements(A, B []int) int {
	ans := 0
	freq := make(map[int]int)
	for _, v := range A {
		freq[v] += 1
	}

	for _, v := range B {
		if freq[v] > 0 {
			ans++
			freq[v]--
		}
	}

	return ans
}

func findThePrefixCommonArray(A []int, B []int) []int {
	var permA []int
	var permB []int
	var ans []int

	for i := range len(A) {
		permA = append(permA, A[i])
		permB = append(permB, B[i])

		x := countCommonElements(permA, permB)
		ans = append(ans, x)
	}

	return ans
}
