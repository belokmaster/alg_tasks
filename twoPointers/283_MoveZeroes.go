package main

func moveZeroes(nums []int) {
	countZeroes := 0
	digitInd := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countZeroes++
		} else {
			nums[digitInd] = nums[i]
			digitInd++
		}
	}

	for i := 0; i < countZeroes; i++ {
		nums[len(nums)-1-i] = 0
	}
}
