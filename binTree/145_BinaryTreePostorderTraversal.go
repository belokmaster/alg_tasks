package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var postOrder func(node *TreeNode)

	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		ans = append(ans, node.Val)
	}

	postOrder(root)
	return ans
}
