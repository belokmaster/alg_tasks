package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans int

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	if leftDepth > rightDepth {
		maxDepth = leftDepth + 1
	} else {
		maxDepth = rightDepth + 1
	}

	if leftDepth+rightDepth > ans {
		ans = leftDepth + rightDepth
	}

	return maxDepth
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	depth(root)
	return ans
}
