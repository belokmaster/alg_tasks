package main

func BTS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	countNodes := 0
	sumNodes := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sumNodes += node.Val
			countNodes++

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return int(sumNodes / countNodes)
}

func averageOfSubtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val == BTS(node) {
				ans++
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ans
}
