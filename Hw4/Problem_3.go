package Hw4

import "fmt"

/*
 * Поиск минимальной глубины бинарного дерева
 *
 * На вход функции подается бинарное дерево. Необходимо найти минимальную
 * глубину дерева. Минимальная глубина — это количество узлов на кратчайшем
 * пути от корневого узла до ближайшего листового узла.
 */

func minHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftH := minHeight(root.Left)
	rightH := minHeight(root.Right)

	if leftH == 0 || rightH == 0 {
		return max(leftH, rightH) + 1
	}

	return min(leftH, rightH) + 1
}

func Problem3() {
	arr := []int{8, 9, 11, 7, 16, 3, 1}
	tree := buildTree(arr, 0)

	fmt.Println("Minimum height of tree is: ", minHeight(tree))
}
