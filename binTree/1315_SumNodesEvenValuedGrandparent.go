package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumEvenGrandparent(root *TreeNode) int {
	ans := 0
	var dfs func(node, parent, grandParent *TreeNode)

	dfs = func(node, parent, grandParent *TreeNode) {
		if node == nil {
			return
		}

		if grandParent != nil && grandParent.Val%2 == 0 {
			ans += node.Val
		}

		dfs(node.Left, node, parent)
		dfs(node.Right, node, parent)
	}

	dfs(root, nil, nil)
	return ans
}
