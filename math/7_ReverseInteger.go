package main

import (
	"math"
)

// Функция divisors разбивает число на цифры и возвращает их в обратном порядке
func divisors1(x int) []int {
	res := []int{}
	for x > 0 {
		res = append(res, x%10)
		x /= 10
	}
	return res
}

func reverse(x int) int {
	flag := 0 // Флаг для отслеживания отрицательных чисел
	if x < 0 {
		flag = 1 // Помечаем, что число отрицательное
		x *= -1  // Преобразуем его в положительное
	}

	div := divisors1(x)
	ans := 0 // Переменная для хранения перевёрнутого числа

	// Составляем перевёрнутое число из массива цифр
	for i := 0; i < len(div); i++ {
		ans += int(math.Pow(10, float64(len(div)-1-i))) * div[i]
	}

	// Проверка на переполнение 32-битного знакового целого числа
	if ans >= int(math.Pow(2, 31))-1 {
		return 0
	}

	// Если изначальное число было отрицательным, возвращаем результат с минусом
	if flag == 1 {
		ans *= -1
	}
	return ans
}

/*
func main() {
	// Тестовые случаи для проверки функции reverse
	testCases := []int{123, -123, 120, 0, int(math.Pow(2, 31)), -int(math.Pow(2, 31))}

	for _, testCase := range testCases {
		fmt.Printf("Original: %d, Reversed: %d\n", testCase, reverse(testCase))
	}
}
*/
