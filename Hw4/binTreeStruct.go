package Hw4

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return [][]string{}
	}

	parent := make(map[*TreeNode]*TreeNode)

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := dfs(node.Left)
		r := dfs(node.Right)

		if node.Left != nil {
			parent[node.Left] = node
		}
		if node.Right != nil {
			parent[node.Right] = node
		}

		return 1 + max(l, r)
	}
	height := dfs(root)

	matrix := make([][]string, height)

	n := int(math.Pow(2, float64(height)) - 1)
	for i := 0; i < height; i++ {
		matrix[i] = make([]string, n)
	}

	type point struct {
		x int
		y int
	}

	positions := make(map[*TreeNode]point)

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		newQueue := []*TreeNode{}

		for _, q := range queue {
			if q == root {
				matrix[0][(n-1)/2] = strconv.Itoa(q.Data)
				positions[q] = point{
					0,
					(n - 1) / 2,
				}
			} else {
				parent_point := positions[parent[q]]
				r, c := parent_point.x, parent_point.y

				offset := int(math.Pow(2, float64(height-r-2)))

				if parent[q].Left == q {
					matrix[r+1][c-offset] = strconv.Itoa(q.Data)
					positions[q] = point{
						r + 1,
						c - offset,
					}
				} else {
					matrix[r+1][c+offset] = strconv.Itoa(q.Data) // Right child
					positions[q] = point{
						r + 1,
						c + offset,
					}
				}
			}
			if q.Left != nil {
				newQueue = append(newQueue, q.Left)
			}
			if q.Right != nil {
				newQueue = append(newQueue, q.Right)
			}
		}

		queue = newQueue
	}

	return matrix
}
