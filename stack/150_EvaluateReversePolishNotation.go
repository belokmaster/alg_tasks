package main

import "strconv"

func oper(a, b, token string) int {
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	if token == "+" {
		return x + y
	}
	if token == "*" {
		return x * y
	}
	if token == "/" {
		return x / y
	}
	if token == "-" {
		return x - y
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := []string{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			result := oper(x, y, tokens[i])
			stack = append(stack, strconv.Itoa(result))
		} else {
			stack = append(stack, tokens[i])
		}
	}
	ans, _ := strconv.Atoi(stack[0])
	return ans
}
