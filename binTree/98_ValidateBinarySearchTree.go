package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveValid(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min || node.Val >= max {
		return false
	}

	return recursiveValid(node.Left, min, node.Val) && recursiveValid(node.Right, node.Val, max)
}

func isValidBST(root *TreeNode) bool {
	return recursiveValid(root, int(math.MinInt64), int(math.MaxInt64))
}
