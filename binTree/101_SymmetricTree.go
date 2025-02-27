package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		levelSize := len(queue)
		nodes := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == nil {
				nodes = append(nodes, -101)
				continue
			}

			nodes = append(nodes, node.Val)
			queue = append(queue, node.Left, node.Right)
		}

		for i := 0; i < len(nodes)/2; i++ {
			if nodes[i] != nodes[len(nodes)-1-i] {
				return false
			}
		}
	}

	return true
}
