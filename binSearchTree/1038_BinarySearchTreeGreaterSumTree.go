package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	sum := 0
	queue := []*TreeNode{}
	node := root

	for len(queue) > 0 || node != nil {
		// до конца правого поддерева
		for node != nil {
			queue = append(queue, node)
			node = node.Right
		}

		node = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		sum += node.Val
		node.Val = sum

		node = node.Left
	}

	return root
}
