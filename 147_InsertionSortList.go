package main

func insertionSortList(head *ListNode) *ListNode {
	var sorted *ListNode // Указатель на отсортированную часть списка
	current := head      // Текущий узел из исходного списка

	for current != nil {
		next := current.Next // Сохраняем следующий узел
		// Вставка в начало отсортированного списка
		if sorted == nil || sorted.Val >= current.Val {
			current.Next = sorted
			sorted = current
		} else {
			// Найти место для вставки
			sortedCurrent := sorted
			for sortedCurrent.Next != nil && sortedCurrent.Next.Val < current.Val {
				sortedCurrent = sortedCurrent.Next
			}
			current.Next = sortedCurrent.Next
			sortedCurrent.Next = current
		}
		current = next // Переходим к следующему узлу
	}

	return sorted
}
