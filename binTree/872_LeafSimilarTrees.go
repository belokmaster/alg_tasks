package main

import "slices"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLeaves(root *TreeNode) []int {
	var leaves []int
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			leaves = append(leaves, node.Val)
			return
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return leaves
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ans1 := getLeaves(root1)
	ans2 := getLeaves(root2)
	return slices.Equal(ans1, ans2)
}
