package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	queue := []*TreeNode{node}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}
	return level
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		leftHeight := height(node.Left)
		rightHeight := height(node.Right)

		if int(math.Abs(float64(leftHeight)-float64(rightHeight))) > 1 {
			return false
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return true
}
