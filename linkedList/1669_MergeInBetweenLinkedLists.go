package main

import "fmt"

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	current := list1
	var nodeA *ListNode
	var nodeB *ListNode
	size := 0

	if a == 0 {
		return list2
	}

	for current != nil {
		size++
		if size == a {
			nodeA = current
		}
		if size == b+2 {
			fmt.Println(size)
			nodeB = current
		}
		current = current.Next
	}

	nodeA.Next = list2
	current = list2
	for current.Next != nil {
		current = current.Next
	}

	current.Next = nodeB

	return list1
}
