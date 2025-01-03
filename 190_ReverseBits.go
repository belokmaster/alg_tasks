package main

import "fmt"

// Функция reverseBits переворачивает биты числа num.
func reverseBits(num uint32) uint32 {
	var result uint32 // Инициализируем переменную для хранения перевёрнутого числа

	// Цикл, который проходит по всем 32 битам числа
	for i := 0; i < 32; i++ {
		result *= 2 // Сдвигаем result влево, умножая на 2, чтобы освободить место для нового бита

		// Если num не равно 0, извлекаем его младший бит
		if num != 0 {
			result += num % 2 // Добавляем младший бит num в result
			num /= 2          // Сдвигаем num вправо, удаляя уже использованный бит
		}
	}

	return result // Возвращаем перевёрнутое число
}

func main() {
	// Тестируем функцию с несколькими числами
	testCases := []uint32{43261596, 4294967293} // Пример чисел для тестирования

	for _, num := range testCases {
		fmt.Printf("Original: %d, Reversed: %d\n", num, reverseBits(num))
	}
}
