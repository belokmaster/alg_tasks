package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// лево корень право
func inorderTraversal(root *TreeNode, nodes *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inorderTraversal(root.Right, nodes)
}

// сортировка
func sortedArrayToBST(nodes []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	root := &TreeNode{Val: nodes[mid]}
	root.Left = sortedArrayToBST(nodes, left, mid-1)
	root.Right = sortedArrayToBST(nodes, mid+1, right)
	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	nodes := []int{}
	inorderTraversal(root, &nodes)
	return sortedArrayToBST(nodes, 0, len(nodes)-1)
}
