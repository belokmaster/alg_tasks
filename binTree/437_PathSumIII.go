package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return countFrom(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func countFrom(node *TreeNode, sum int) int {
	if node == nil {
		return 0
	}

	ans := 0
	if node.Val == sum {
		ans++
	}

	ans += countFrom(node.Left, sum-node.Val)
	ans += countFrom(node.Right, sum-node.Val)

	return ans
}
