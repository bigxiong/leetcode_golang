package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var min = math.MaxInt32
	if root.Left != nil {
		leftDepth := minDepth(root.Left)
		if leftDepth < min {
			min = leftDepth
		}
	}
	if root.Right != nil {
		rightDepth := minDepth(root.Right)
		if rightDepth < min {
			min = rightDepth
		}
	}

	return min + 1
}

// BFS
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var depth int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			// poll queue head
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

func minDepth3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := math.MaxInt32
	left := minDepth3(root.Left)
	minDepth = min(minDepth, left)

	right := minDepth3(root.Right)
	minDepth = min(minDepth, right)

	return minDepth + 1
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {

}
