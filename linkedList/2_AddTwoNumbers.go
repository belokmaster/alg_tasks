package main

/*  Вам даны два непустых связанных списка, представляющих два неотрицательных целых числа.
    Цифры хранятся в обратном порядке, и каждый из их узлов содержит одну цифру.
    Сложите два числа и верните сумму в виде связанного списка.
    Вы можете предположить, что эти два числа не содержат никакого начального нуля, за исключением самого числа 0.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers складывает два числа, представленные двумя связанными списками, и возвращает сумму в виде связанного списка.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Создаем вспомогательный головной узел
	dummyHead := &ListNode{Val: 0}
	current := dummyHead // Инициализируем текущий узел как вспомогательный головной узел
	carry := 0           // Инициализируем перенос равным 0

	// Цикл продолжается, пока не пройдем все узлы обоих списков и не будет переноса
	for l1 != nil || l2 != nil || carry != 0 {
		x, y := 0, 0 // Инициализируем значения для текущих узлов l1 и l2
		if l1 != nil {
			x = l1.Val   // Получаем значение из l1, если оно существует
			l1 = l1.Next // Переходим к следующему узлу в l1
		}
		if l2 != nil {
			y = l2.Val   // Получаем значение из l2, если оно существует
			l2 = l2.Next // Переходим к следующему узлу в l2
		}

		sum := carry + x + y                    // Суммируем значения текущих узлов и перенос
		carry = sum / 10                        // Вычисляем новый перенос
		current.Next = &ListNode{Val: sum % 10} // Создаем новый узел для результата и добавляем его в конец списка
		current = current.Next                  // Переходим к следующему узлу результата
	}

	return dummyHead.Next // Возвращаем результирующий список, пропуская вспомогательный головной узел
}
