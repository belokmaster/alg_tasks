package main

// inorderTraversal выполняет обход бинарного дерева в порядке "in-order" (левый узел -> текущий узел -> правый узел)
func inorderTraversal(root *TreeNode) []int {
	// Базовый случай: если узел пустой, возвращаем пустой массив
	if root == nil {
		return []int{}
	}

	// Рекурсивно обходим левое поддерево и сохраняем его элементы
	left := inorderTraversal(root.Left)
	// Добавляем значение текущего узла в массив
	left = append(left, root.Val)

	// Рекурсивно обходим правое поддерево и сохраняем его элементы
	right := inorderTraversal(root.Right)
	// Добавляем элементы из правого поддерева в массив, объединяя их с левыми элементами и текущим узлом
	left = append(left, right...)

	return left
}
