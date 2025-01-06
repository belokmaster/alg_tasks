package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func maxDepthIterative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Очередь, которая будет хранить узлы текущего уровня.
	queue := []*TreeNode{root}
	depth := 0

	// Пока есть узлы в очереди
	for len(queue) > 0 {
		// Обрабатываем текущий уровень
		levelSize := len(queue) // Количество узлов на текущем уровне
		for i := 0; i < levelSize; i++ {
			// Удаляем узел из очереди
			node := queue[0]
			queue = queue[1:]

			// Добавляем дочерние узлы в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// Увеличиваем глубину после обработки уровня
		depth++
	}

	return depth
}

// DFS
func maxDepth2(root *TreeNode) int {
	// Базовый случай: если узел пустой, глубина равна 0
	if root == nil {
		return 0
	}

	// Рекурсивно вычисляем глубину левого поддерева
	leftDepth := maxDepth2(root.Left)
	// Рекурсивно вычисляем глубину правого поддерева
	rightDepth := maxDepth2(root.Right)

	// Сравниваем глубину левого и правого поддеревьев
	// Глубина текущего узла = 1 (текущий уровень) + максимальная глубина поддеревьев
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
