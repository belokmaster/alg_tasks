package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	digits := make(map[int]bool)
	current := head

	sortedDigits := []int{}
	for current != nil {
		if digits[current.Val] {
			current = current.Next
			continue
		}
		digits[current.Val] = true
		sortedDigits = append(sortedDigits, current.Val)
		current = current.Next
	}

	sort.Ints(sortedDigits)
	ans := &ListNode{}
	head = ans
	current = head

	for i := 0; i < len(sortedDigits); i++ {
		newNode := &ListNode{Val: sortedDigits[i]}

		for current.Next != nil {
			current = current.Next
		}

		current.Next = newNode
	}

	return head.Next
}
