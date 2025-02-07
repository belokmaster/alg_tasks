package main

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	if digits[len(digits)-1] != 10 {
		return digits
	}

	newDigit := append([]int{0}, digits...)
	addFlag := 0
	for i := 0; i < len(newDigit); i++ {
		pos := len(newDigit) - i - 1
		if addFlag == 1 {
			newDigit[pos]++
			addFlag = 0
		}
		if newDigit[pos] > 9 {
			newDigit[pos] = 0
			addFlag = 1
		}
	}

	if newDigit[0] == 0 {
		return newDigit[1:]
	}
	return newDigit
}
