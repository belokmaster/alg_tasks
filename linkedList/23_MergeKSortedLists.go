package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var result *ListNode

	if left.Val < right.Val {
		result = &ListNode{Val: left.Val}
		result.Next = merge(left.Next, right)
	} else {
		result = &ListNode{Val: right.Val}
		result.Next = merge(left, right.Next)
	}

	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	ans := lists[0]
	for i := 1; i < len(lists); i++ {
		ans = merge(ans, lists[i])
	}

	return ans
}
