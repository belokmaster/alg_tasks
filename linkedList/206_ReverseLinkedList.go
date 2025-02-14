package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var newList *ListNode // указатель на новый перевёрнутый список.
	current := head       // указатель на голову списка

	for current != nil {
		next := current.Next
		current.Next = newList
		newList = current
		current = next
	}

	return newList
}
