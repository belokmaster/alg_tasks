package main

func localMax(arr []int) int {
	ans := arr[0]
	localMaxInd := 0
	for i, num := range arr {
		if num > ans {
			ans = num
			localMaxInd = i
		}
	}
	return localMaxInd
}

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	localMaxInd := localMax(arr)
	if localMaxInd == len(arr)-1 || localMaxInd == 0 {
		return false
	}

	for i := 0; i < localMaxInd-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}

	for i := localMaxInd + 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}
