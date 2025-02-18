package main

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	digits := make(map[int]int)

	current := head
	for current != nil {
		digits[current.Val]++
		current = current.Next
	}

	sortedDigits := []int{}
	for key, values := range digits {
		if values == 1 {
			sortedDigits = append(sortedDigits, key)
		}
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
