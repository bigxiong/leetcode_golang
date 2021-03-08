package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var hl, hr int
	l, r := root, root
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	if hl == hr {
		return int(math.Pow(float64(2), float64(hl))) - 1
	}

	return  1 + countNodes(root.Left) + countNodes(root.Right)
}

func main() {
	
}
