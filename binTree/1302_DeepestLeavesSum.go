package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countLevels(root *TreeNode) int {
	level := 0
	if root == nil {
		return level
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)

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

		level++
	}

	return level
}

func deepestLeavesSum(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}

	level := 0
	deepLevel := countLevels(root)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if level == deepLevel-1 {
				ans += node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}
	return ans
}
