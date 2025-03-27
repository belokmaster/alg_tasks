package main

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{}

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[index] = i - index
		}
		stack = append(stack, i)
	}

	return ans
}
