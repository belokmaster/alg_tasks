package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTreeTheSame(root *TreeNode, subRoot *TreeNode) bool {
	queue := []*TreeNode{root}
	subQueue := []*TreeNode{subRoot}

	for len(queue) > 0 {
		if len(subQueue) != len(queue) {
			return false
		}

		levelSize := len(queue)
		levelSizeSub := len(subQueue)

		if levelSize != levelSizeSub {
			return false
		}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			subNode := subQueue[0]
			subQueue = subQueue[1:]

			if subNode.Val != node.Val {
				return false
			}

			if (node.Left == nil) != (subNode.Left == nil) {
				return false
			}
			if node.Left != nil && subNode.Left != nil {
				queue = append(queue, node.Left)
				subQueue = append(subQueue, subNode.Left)
			}

			if (node.Right == nil) != (subNode.Right == nil) {
				return false
			}
			if node.Right != nil && subNode.Right != nil {
				queue = append(queue, node.Right)
				subQueue = append(subQueue, subNode.Right)
			}
		}
	}

	return true
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if checkTreeTheSame(node, subRoot) {
				return true
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return false
}
