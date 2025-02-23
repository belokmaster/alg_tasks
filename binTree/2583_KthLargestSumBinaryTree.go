package main

import (
	"slices"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sum(arr []int) int {
	ans := 0
	for _, num := range arr {
		ans += num
	}
	return ans
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return 0
	}

	finalSum := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		sumLevel := []int{}
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sumLevel = append(sumLevel, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		finalSum = append(finalSum, sum(sumLevel))
	}

	sort.Ints(finalSum)
	slices.Reverse(finalSum)
	if k > len(finalSum) {
		return -1
	}
	return int64(finalSum[k-1])
}
