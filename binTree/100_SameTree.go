package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bfs(root *TreeNode) []string {
	queue := []*TreeNode{root}
	nodes := []string{}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			nodes = append(nodes, strconv.Itoa(node.Val))

			if node.Left != nil {
				queue = append(queue, node.Left)
			} else {
				nodes = append(nodes, "L")
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			} else {
				nodes = append(nodes, "R")
			}
		}
	}

	return nodes
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	}

	if q == nil {
		return false
	}

	firstNodes := bfs(p)
	secondNodes := bfs(q)

	for i := range firstNodes {
		if firstNodes[i] != secondNodes[i] {
			return false
		}
	}

	return true
}
