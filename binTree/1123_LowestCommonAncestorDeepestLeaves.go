package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findDeepestLeaves(node *TreeNode, depth int) (int, *TreeNode) {
	if node == nil {
		return depth, nil
	}

	leftDepth, leftLCA := findDeepestLeaves(node.Left, depth+1)
	rightDepth, rightLCA := findDeepestLeaves(node.Right, depth+1)

	if leftDepth == rightDepth {
		return leftDepth, node
	}

	if leftDepth > rightDepth {
		return leftDepth, leftLCA
	}
	return rightDepth, rightLCA
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := findDeepestLeaves(root, 0)
	return lca
}
