package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(list *ListNode) *ListNode {
	var prev *ListNode
	current := list

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head = reverse(head)

	if n == 1 {
		head = head.Next
		return reverse(head)
	}

	current := head
	for i := 0; i < n-2 && current != nil; i++ {
		current = current.Next
	}

	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}

	return reverse(head)
}
