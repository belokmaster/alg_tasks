package main

// DFS
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

// BFS
func rangeSumBST1(root *TreeNode, low int, high int) int {
	ans := 0
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val >= low && node.Val <= high {
			ans += node.Val
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return ans
}
