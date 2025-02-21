package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	high := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		if high%2 == 1 {
			left := 0
			right := size - 1

			for left < right {
				queue[left].Val, queue[right].Val = queue[right].Val, queue[left].Val
				left++
				right--
			}
		}

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		high++
	}

	return root
}
