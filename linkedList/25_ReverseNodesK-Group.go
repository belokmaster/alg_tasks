package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse1(head *ListNode) *ListNode {
	var newList *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = newList
		newList = current
		current = next
	}

	return newList
}

func insertAtEnd(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}
	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var newNode *ListNode
	var answer *ListNode
	count := 0
	current := head

	for current != nil {
		newNode = insertAtEnd(newNode, current.Val)
		count++
		if count%k == 0 {
			newNode = reverse1(newNode)
			if answer == nil {
				answer = newNode
			} else {
				tail := answer
				for tail.Next != nil {
					tail = tail.Next
				}
				tail.Next = newNode
			}

			newNode = nil
			count = 0
		}
		current = current.Next
	}

	if answer == nil {
		answer = newNode
	} else {
		tail := answer
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = newNode
	}

	return answer
}
