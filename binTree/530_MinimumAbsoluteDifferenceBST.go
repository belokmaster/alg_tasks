package main

import (
	"math"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	nodes := []int{}
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			nodes = append(nodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	// можно без сортировки делать preorder
	sort.Ints(nodes)
	minDifference := int(math.MaxInt64)
	for i := 1; i < len(nodes); i++ {
		if nodes[i]-nodes[i-1] < minDifference {
			minDifference = nodes[i] - nodes[i-1]
		}
	}

	return minDifference
}
