package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	beforePivot := &ListNode{}
	afterPivot := &ListNode{}

	beforePivotCurrent := beforePivot
	afterPivotCurrent := afterPivot

	current := head
	for current != nil {
		newNode := &ListNode{Val: current.Val}
		if current.Val < x {
			beforePivotCurrent.Next = newNode
			beforePivotCurrent = newNode
		} else {
			afterPivotCurrent.Next = newNode
			afterPivotCurrent = newNode
		}
		current = current.Next
	}

	afterPivotCurrent.Next = nil
	beforePivotCurrent.Next = afterPivot.Next

	return beforePivot.Next
}
