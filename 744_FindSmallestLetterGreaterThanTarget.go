package main

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	left := 0
	right := n - 1
	ans := -1

	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] < target {
			left = mid + 1
		} else if letters[mid] > target {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if ans == -1 {
		return letters[0]
	}

	return letters[ans]
}
