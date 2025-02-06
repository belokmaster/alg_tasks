package main

func maxElement(arr []int, ind int) int {
	ans := arr[ind+1]
	for _, num := range arr[ind+1:] {
		if num > ans {
			ans = num
		}
	}
	return ans
}

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = maxElement(arr, i)
	}
	arr[len(arr)-1] = -1
	return arr
}
