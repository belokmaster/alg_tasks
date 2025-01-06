package main

// isValid проверяет, является ли строка из скобок допустимой
func isValid(s string) bool {
	// Создаем пустой стек для хранения ожидаемых закрывающих скобок
	stack := make([]rune, 0)

	m := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, c := range s {
		// Если символ — открывающая скобка
		if char, ok := m[c]; ok {
			// Добавляем соответствующую закрывающую скобку в стек
			stack = append(stack, char)
		} else {
			// Если стек пуст или верхняя скобка не совпадает с текущей закрывающей
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			// Удаляем верхний элемент из стека, так как скобки совпали
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

/*
func main() {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"",
		"{{{{",
		"(([]){})",
	}

	for _, testCase := range testCases {
		fmt.Printf("Строка: %q -> Допустимо: %v\n", testCase, isValid(testCase))
	}
}
*/
