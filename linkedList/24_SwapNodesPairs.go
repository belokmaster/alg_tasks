package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next
	newHead := p2

	for p2 != nil {
		next1 := p2.Next
		var next2 *ListNode

		if next1 != nil {
			next2 = next1.Next
		}

		p2.Next = p1
		if next2 == nil {
			p1.Next = next1
		} else {
			p1.Next = next2
		}

		p1 = next1
		p2 = next2
	}

	return newHead
}

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next.Next
	head.Next.Next = head
	newHead := head.Next
	head.Next = swapPairsRecursive(next)

	return newHead
}
