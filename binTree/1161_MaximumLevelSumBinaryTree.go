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
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := int(math.MinInt64)
	level := 0
	ansLevel := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		tempSum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			tempSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
		if tempSum > ans {
			ans = tempSum
			ansLevel = level
		}
	}

	return ansLevel
}
