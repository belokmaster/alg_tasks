package main

// кол-во способов подняться на лестницу - рекуррентное соотношение совпадающее с посл. фибоначчи
func climbStairs(n int) int {
	a, b := 1, 1

	for n > 1 {
		a, b = b, a+b
		n--
	}

	return b
}
