package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfsGood(node *TreeNode, maxFar int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= maxFar {
		count = 1
		maxFar = node.Val
	}

	return count + dfsGood(node.Left, maxFar) + dfsGood(node.Right, maxFar)
}

func goodNodes(root *TreeNode) int {
	return dfsGood(root, root.Val)
}
