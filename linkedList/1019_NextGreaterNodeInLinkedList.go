package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findMaxElement(head *ListNode) int {
	current := head
	maxElement := head.Val
	current = current.Next

	for current != nil {
		if current.Val > maxElement {
			maxElement = current.Val
			return maxElement
		}
		current = current.Next
	}

	if maxElement == head.Val {
		return 0
	}
	return maxElement
}

func nextLargerNodes(head *ListNode) []int {
	ans := []int{}
	current := head

	for current != nil {
		x := findMaxElement(current)
		ans = append(ans, x)
		current = current.Next
	}

	return ans
}
