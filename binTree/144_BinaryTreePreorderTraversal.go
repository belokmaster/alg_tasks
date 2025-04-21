package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var preorderTraversal func(node *TreeNode)

	preorderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorderTraversal(node.Left)
		preorderTraversal(node.Right)
	}

	preorderTraversal(root)
	return ans
}
