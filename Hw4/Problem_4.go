package Hw4

import "fmt"

/*
 * Поиск произведения максимального и минимального элементов
 */

func minMaxProduct(root *TreeNode) int {
	minNode := root
	maxNode := root

	for minNode.Left != nil {
		minNode = minNode.Left
	}

	for maxNode.Right != nil {
		maxNode = maxNode.Right
	}

	return minNode.Data * maxNode.Data
}

func Problem4() {
	arr := []int{8, 9, 11, 7, 16, 3, 1}
	tree := buildTree(arr, 0)

	fmt.Println("Product of minimum and maximum values is: ", minMaxProduct(tree))
}
