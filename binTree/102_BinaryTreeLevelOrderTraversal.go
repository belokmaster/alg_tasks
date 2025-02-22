package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sizeLevel := len(queue)
		nodeArr := []int{}

		for i := 0; i < sizeLevel; i++ {
			node := queue[0]
			queue = queue[1:]
			nodeArr = append(nodeArr, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		ans = append(ans, nodeArr)
	}

	return ans
}
