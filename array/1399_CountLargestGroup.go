package main

func sumOfDigits(n int) int {
	ans := 0
	for n > 0 {
		ans += n % 10
		n /= 10
	}
	return ans
}

func countLargestGroup(n int) int {
	hashMap := make(map[int][]int)
	for i := 1; i <= n; i++ {
		size := sumOfDigits(i)
		hashMap[size] = append(hashMap[size], i)
	}

	ansLen := 0
	ans := 0
	for _, value := range hashMap {
		if len(value) > ansLen {
			ansLen = len(value)
			ans = 1
		} else if len(value) == ansLen {
			ans++
		}
	}
	return ans
}
